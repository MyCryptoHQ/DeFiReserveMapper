resource "aws_cloudwatch_event_rule" "defi_updater_cloudwatch_rule" {
  name                = "defi-updater"
  description         = "Cloudwatch rule to trigger defi updater task."
  schedule_expression = var.schedule
}

resource "aws_cloudwatch_event_target" "defi_updater_cloudwatch_target" {
  arn  = aws_lambda_function.defi-mapper-updater-lambda.arn
  rule = aws_cloudwatch_event_rule.defi_updater_cloudwatch_rule.name
}

resource "aws_lambda_permission" "allow_cloudwatch_trigger_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.defi-mapper-updater-lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.defi_updater_cloudwatch_rule.arn
}
