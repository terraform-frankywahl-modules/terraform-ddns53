output "url" {
  description = "The URL where the endpoint can be invoked"
  value = aws_api_gateway_deployment.deployment.invoke_url
}

output "username" {
  description = "The username to use for basic authentication of the endpoint"
  value = local.username
}

output "password" {
  description = "The password to use for basic authentication of the endpoint"
  value = local.password
  sensitive = true
}
