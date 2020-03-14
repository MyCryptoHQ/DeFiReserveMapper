resource "aws_s3_bucket" "defi_s3_bucket" {
  bucket = "defi-reserve-mapper"
  acl    = "private"

  versioning {
    enabled = true
  }

  tags = {
    Name        = "app"
    Environment = "defi-reserve-mapper"
  }
}
