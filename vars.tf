variable "domain" {
  description = "The TLD of the domain we want to create a DDNS for"
  type        = string
}

variable "subdomain" {
  description = "The subdomain that needs to be updated with DDNS"
  type        = string
  default     = "home"
}

variable "enable_cloudwatch" {
  description = "Enable Lambda to log cloudwatch events"
  type        = bool
  default     = true
}

variable "tags" {
  description = "A mapping of tags to assign to the resources created"
  type        = map(string)
  default     = {}
}

variable "username" {
  description = "Username for the basic authentication. One will be created if none are given"
  type        = string
  default     = null
}

variable "password" {
  description = "Password for the basic authentication. One will be created if none are given"
  type        = string
  default     = null
}
