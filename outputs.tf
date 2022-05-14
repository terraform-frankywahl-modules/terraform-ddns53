output "url" {
  value = aws_api_gateway_deployment.deployment.invoke_url
}

output "username" {
  value = local.username
}

output "password" {
  value = local.password
  sensitive = true
}
