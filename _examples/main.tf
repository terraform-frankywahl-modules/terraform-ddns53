module "ddns" {
  //source            = "git@github.com:terraform-frankywahl-modules/terraform-ddns53.git"
  source            = "../."
  domain            = var.domain
  subdomain         = var.subdomain
  enable_cloudwatch = false
}
