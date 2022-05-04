data "aws_route53_zone" "domain" {
  name = var.domain
}

resource "aws_route53_record" "ddns" {
  zone_id = data.aws_route53_zone.domain.id
  name    = "${var.subdomain}.${var.domain}"
  type    = "A"
  ttl     = 60
  records = ["0.0.0.0"]

  lifecycle {
    ignore_changes = [
      records
    ]
  }
}
