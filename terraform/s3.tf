resource "aws_s3_bucket" "defi_s3_bucket" { // creates a new `aws_s3_bucket` resource and gives it an id `defi_s3_bucket`
  bucket        = var.bucket                // names the bucket `defi-reserve-mapper`
  acl           = "public-read"             // defines the bucket as private
  force_destroy = true

  website {
    index_document = "outputFile.json"
    error_document = "tmp/outputFile.json"
  }

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET"]
    allowed_origins = ["*"]
    max_age_seconds = 1800
  }

  versioning {
    enabled = true // enables versioning for the `defi_s3_bucket` resource
  }
  policy = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [
    {
      "Sid": "PublicReadForGetBucketObjects",
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::${var.bucket}/*"
    }
  ]
}
EOF

  tags = {
    Name = "defi-reserve-mapper" // sets bucket tag in aws to "defi-reserve-mapper"
  }
}
