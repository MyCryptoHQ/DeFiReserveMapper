terraform {
  backend "s3" {
    bucket = "defi-reserve-mapper-tf-prd"
    key    = "terraform.tfstate"
    region = "us-east-1"
  }
}
