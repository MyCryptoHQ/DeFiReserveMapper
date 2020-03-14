terraform {
  backend "s3" {
    bucket = "defi-reserve-mapper-tf"
    key    = "terraform.tfstate"
    region = "us-east-1"
  }
}
