# Dynamic DNS & AWS

Dynamic DNS API usage with AWS Lambda

## Usage

```tf
module "ddns" {
  source            = "git@github.com:terraform-frankywahl-modules/terraform-ddns53.git"
  domain            = "example.com" # Assumes you alread created the hosted zone with this domain
  subdomain         = "home"
  enable_cloudwatch = true
}
```

This creates a Route53 record (`home.example.com`), an API Gateway and Lambda that that will allow updating that record.

### Notes

* API Gateway acts as a full proxy
* Username and Password are stored as environment variables on the lambda. Please use those when configuring the service with DynDNS
* Since there is no way of knowing if the binary was changed or not, the go binary will be recreated on every deploy

## Requirements

* [Go](https://go.dev/) - for compiling the soure code
* [Terraform](https://www.terraform.io/)
* [AWS](https://console.aws.amazon.com/console/home) account with an existent hosted zone

### References

References that were used to help create this module:

* https://aws.amazon.com/blogs/startups/how-to-build-a-serverless-dynamic-dns-system-with-aws/
  * https://github.com/awslabs/route53-dynamic-dns-with-lambda
* https://github.com/sTywin/dyndns53
* https://github.com/jcmturner/ddns

## Providers

| Name | Version |
|------|---------|
| <a name="provider_archive"></a> [archive](#provider\_archive) | 2.2.0 |
| <a name="provider_aws"></a> [aws](#provider\_aws) | 4.13.0 |
| <a name="provider_null"></a> [null](#provider\_null) | 3.1.1 |
| <a name="provider_random"></a> [random](#provider\_random) | 3.1.3 |

## Resources

| Name | Type |
|------|------|
| [aws_api_gateway_deployment.deployment](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_deployment) | resource |
| [aws_api_gateway_integration.options_integration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_integration) | resource |
| [aws_api_gateway_integration.request_integration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_integration) | resource |
| [aws_api_gateway_integration.request_integration_root](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_integration) | resource |
| [aws_api_gateway_integration_response.options_integration_response](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_integration_response) | resource |
| [aws_api_gateway_method.options_method](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method) | resource |
| [aws_api_gateway_method.request_method](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method) | resource |
| [aws_api_gateway_method.request_method_root](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method) | resource |
| [aws_api_gateway_method_response.options_200](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method_response) | resource |
| [aws_api_gateway_resource.proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_resource) | resource |
| [aws_api_gateway_rest_api.api](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api) | resource |
| [aws_iam_policy.lambda_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy) | resource |
| [aws_iam_role.lambda_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy.lambda](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy_attachment.cloudwatch](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_lambda_function.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function) | resource |
| [aws_lambda_permission.allow_api_gateway](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission) | resource |
| [aws_route53_record.ddns](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record) | resource |
| [null_resource.artifact](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |
| [random_password.password](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password) | resource |
| [random_string.username](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string) | resource |
| [archive_file.artifact](https://registry.terraform.io/providers/hashicorp/archive/latest/docs/data-sources/file) | data source |
| [aws_iam_policy.cloudwatch](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy) | data source |
| [aws_route53_zone.domain](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_domain"></a> [domain](#input\_domain) | The TLD of the domain we want to create a DDNS for | `string` | n/a | yes |
| <a name="input_enable_cloudwatch"></a> [enable\_cloudwatch](#input\_enable\_cloudwatch) | Enable Lambda to log cloudwatch events | `bool` | `true` | no |
| <a name="input_password"></a> [password](#input\_password) | Password for the basic authentication. One will be created if none are given | `string` | `null` | no |
| <a name="input_subdomain"></a> [subdomain](#input\_subdomain) | The subdomain that needs to be updated with DDNS | `string` | `"home"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A mapping of tags to assign to the resources created | `map(string)` | `{}` | no |
| <a name="input_username"></a> [username](#input\_username) | Username for the basic authentication. One will be created if none are given | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_password"></a> [password](#output\_password) | The password to use for basic authentication of the endpoint |
| <a name="output_url"></a> [url](#output\_url) | The URL where the endpoint can be invoked |
| <a name="output_username"></a> [username](#output\_username) | The username to use for basic authentication of the endpoint |
