resource "aws_lambda_function" "defi-mapper-updater-lambda" {
  filename      = "../app.zip"
  function_name = "defi-mapper-updater-lambda"
  role          = aws_iam_role.defi-mapper-lambda-role.arn
  handler       = "bin/app"
  timeout       = 600

  source_code_hash = filebase64sha256("../app.zip")

  runtime = "go1.x"

  environment {
    variables = {
      region = var.region
      bucket = var.bucket
    }
  }
}

resource "aws_iam_policy" "lambda_logging" {
  name        = "lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:log-group:/aws/lambda/*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role" "defi-mapper-lambda-role" {
  name = "defi-mapper-lambda-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}
resource "aws_iam_policy" "defi-lambda-s3-policy" {
  name = "defi-lambda-s3-policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
      {
          "Effect": "Allow",
          "Action": [
            "s3:*"
          ],
          "Resource": [
            "arn:aws:s3:::${var.bucket}",
            "arn:aws:s3:::${var.bucket}/*"
          ]
      }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.defi-mapper-lambda-role.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

resource "aws_iam_role_policy_attachment" "defi-s3-lambda" {
  role       = aws_iam_role.defi-mapper-lambda-role.name
  policy_arn = aws_iam_policy.defi-lambda-s3-policy.arn

}
