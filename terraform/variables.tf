variable "region" {
  type        = string
  description = "The region to operate out of."
  default     = "us-east-1"
}

variable "profile" {
  type = string
  description = "Name of AWS credentials profile."
  default = "default"
}
