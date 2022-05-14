/**
 * # Dynamic DNS & AWS
 *
 * Dynamic DNS API usage with AWS Lambda
 *
 * ## Usage
 *
 * ```tf
 * module "ddns" {
 *   source            = "git@github.com:terraform-frankywahl-modules/terraform-ddns53.git"
 *   domain            = "example.com" # Assumes you alread created the hosted zone with this domain
 *   subdomain         = "home"
 *   enable_cloudwatch = true
 * }
 * ```
 *
 * This creates a Route53 record (`home.example.com`), an API Gateway and Lambda that that will allow updating that record.
 *
 * ### Notes
 *
 * * API Gateway acts as a full proxy
 * * Username and Password are stored as environment variables on the lambda. Please use those when configuring the service with DynDNS
 * * Since there is no way of knowing if the binary was changed or not, the go binary will be recreated on every deploy
 *
 * ## Requirements
 *
 * * [Go](https://go.dev/) - for compiling the soure code
 * * [Terraform](https://www.terraform.io/)
 * * [AWS](https://console.aws.amazon.com/console/home) account with an existent hosted zone
 *
 * ### References
 *
 * References that were used to help create this module:
 *
 * * https://aws.amazon.com/blogs/startups/how-to-build-a-serverless-dynamic-dns-system-with-aws/
 *   * https://github.com/awslabs/route53-dynamic-dns-with-lambda
 * * https://github.com/sTywin/dyndns53
 * * https://github.com/jcmturner/ddns
 */
