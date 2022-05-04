locals {
  name = replace("ddns-${var.domain}", ".", "-")
}
