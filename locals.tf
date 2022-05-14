locals {
  name = replace("ddns-${var.domain}", ".", "-")

  username = var.username == null ? random_string.username[0].result : var.username
  password = var.password == null ? random_password.password[0].result : var.password
}
