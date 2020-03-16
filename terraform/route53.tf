data "aws_route53_zone" "mycryptoapi_hosted_zone" {
  name         = "mycryptoapi.com."
  private_zone = false
}

resource "aws_route53_record" "defi_mycryptoapi_record" {
  zone_id = data.aws_route53_zone.mycryptoapi_hosted_zone.zone_id
  name    = var.bucket
  type    = "A"

  depends_on = [aws_cloudfront_distribution.s3_distribution]

  alias {
    evaluate_target_health = false
    name                   = aws_cloudfront_distribution.s3_distribution.domain_name
    zone_id                = aws_cloudfront_distribution.s3_distribution.hosted_zone_id
  }
}
