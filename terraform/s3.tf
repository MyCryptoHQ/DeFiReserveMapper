resource "aws_s3_bucket" "defi_s3_bucket" { // creates a new `aws_s3_bucket` resource and gives it an id `defi_s3_bucket`
  bucket = "defi-reserve-mapper"            // names the bucket `defi-reserve-mapper`
  acl    = "private"                        // defines the bucket as private

  versioning {
    enabled = true // enables versioning for the `defi_s3_bucket` resource
  }

  tags = {
    Name        = "defi-reserve-mapper" // sets bucket tag in aws to "defi-reserve-mapper"
    Environment = "defi-reserve-mapper" // sets environment that the `defi-reserve-mapper` bucket belongs to
  }
}
