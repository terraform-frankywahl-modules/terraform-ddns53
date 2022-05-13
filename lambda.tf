resource "null_resource" "artifact" {
  triggers = {
    always_run = timestamp()
  }
  provisioner "local-exec" {
    working_dir = "${path.module}/code/"
    command     = "make lambda"
  }
}
data "archive_file" "artifact" {
  type        = "zip"
  source_file = "${path.module}/code/lambda"
  output_path = "${path.module}/code/lambda.zip"

  depends_on = [null_resource.artifact]
}

resource "aws_iam_role" "lambda_role" {
  name = "DDNS-${var.domain}-lambda-role"
  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : "sts:AssumeRole",
        "Principal" : {
          "Service" : "lambda.amazonaws.com"
        },
        "Effect" : "Allow",
      }
    ]
  })
  path = "/service-role/"

  tags = merge({
    terraform = true
  }, var.tags)
}

resource "aws_lambda_function" "main" {
  function_name = local.name
  handler       = "lambda"
  role          = aws_iam_role.lambda_role.arn
  memory_size   = 128

  runtime          = "go1.x"
  filename         = "${path.module}/code/lambda.zip"
  source_code_hash = data.archive_file.artifact.output_base64sha256

  tags = merge({
    terraform = true
  }, var.tags)

  environment {
    variables = {
      "USERNAME" = local.username
      "PASSWORD" = local.password
      "FQDN"     = resource.aws_route53_record.ddns.name
      "ZONE_ID"  = data.aws_route53_zone.domain.zone_id
    }
  }

  provisioner "local-exec" {
    working_dir = "${path.module}/code"
    command     = "make clean"
  }

  depends_on = [
    null_resource.artifact
  ]
}

resource "aws_iam_policy" "lambda_policy" {
  name = local.name
  policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [{
        "Effect" : "Allow",
        "Action" : [
          "route53:ChangeResourceRecordSets"
        ],
        "Resource" : "arn:aws:route53:::hostedzone/${data.aws_route53_zone.domain.zone_id}"
        }, {
        "Effect" : "Allow",
        "Action" : [
          "route53:ListResourceRecordSets"
        ],
        "Resource" : "arn:aws:route53:::hostedzone/${data.aws_route53_zone.domain.zone_id}"
        }, {
        "Effect" : "Allow",
        "Action" : [
          "route53:GetChange"
        ],
        "Resource" : "arn:aws:route53:::change/*"
        }, {
        "Effect" : "Allow",
        "Action" : [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource" : "arn:aws:logs:*:*:*"
      }]
    }
  )
}

resource "aws_iam_role_policy" "lambda" {
  name = "DDNS-Update-Policy-${local.name}"
  role = aws_iam_role.lambda_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "route53:ChangeResourceRecordSets"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:route53:::hostedzone/${data.aws_route53_zone.domain.id}"
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "cloudwatch" {
  count      = var.enable_cloudwatch ? 1 : 0
  role       = aws_iam_role.lambda_role.id
  policy_arn = data.aws_iam_policy.cloudwatch.arn
}

data "aws_iam_policy" "cloudwatch" {
  name = "AWSLambdaBasicExecutionRole"
}

resource "random_string" "username" {
  count  = var.username == null ? 1 : 0
  length = 16
}

resource "random_password" "password" {
  count  = var.password == null ? 1 : 0
  length = 16
}

resource "aws_lambda_function_url" "endpoint" {
  function_name      = aws_lambda_function.main.function_name
  authorization_type = "NONE"
}
