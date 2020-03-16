variable "region" {
  type        = string
  description = "The region to operate out of."
}

variable "profile" {
  type        = string
  description = "Name of AWS credentials profile."
}

variable "bucket" {
  type        = string
  description = "Name of defi-reserve-mapper S3 bucket."
}

variable "schedule" {
  type        = string
  description = "Schedule expression for Cloudwatch event trigger."
}
