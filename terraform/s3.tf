resource "aws_s3_bucket" "defi_s3_bucket" { // creates a new `aws_s3_bucket` resource and gives it an id `defi_s3_bucket`
  bucket        = var.bucket                // names the bucket `defi-reserve-mapper`
  acl           = "private"                 // defines the bucket as private
  force_destroy = true

  versioning {
    enabled = true // enables versioning for the `defi_s3_bucket` resource
  }

  tags = {
    Name = "defi-reserve-mapper" // sets bucket tag in aws to "defi-reserve-mapper"
  }
}
