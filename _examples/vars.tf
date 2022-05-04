variable "domain" {
  description = "The TLD of the domain we want to create a DDNS for"
  type        = string
  default     = "example.com"
}

variable "subdomain" {
  description = "The subdomain that needs to be updated with DDNS"
  type        = string
  default     = "test"
}
