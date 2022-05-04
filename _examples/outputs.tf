output "url" {
  value = module.ddns.url
}

output "username" {
  value = module.ddns.username
}

output "password" {
  value     = module.ddns.password
  sensitive = true
}
