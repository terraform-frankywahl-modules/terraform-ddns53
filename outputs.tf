output "url" {
  value = aws_api_gateway_deployment.deployment.invoke_url
}

output "username" {
  value = random_string.username.result
}

output "password" {
  value = random_password.password.result
}
