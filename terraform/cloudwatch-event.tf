resource "aws_cloudwatch_event_rule" "defi_updater_cloudwatch_rule" {
  name        = "defi-updater"
  description = "Cloudwatch rule to trigger defi updater task."

  schedule_expression = "rate(5 minutes)"
}

// TODO once we know where the updater will be running
/*
resource "aws_cloudwatch_event_target" "defi_updater_cloudwatch_target" {
  arn = ""
  rule = aws_cloudwatch_event_rule.defi_updater_cloudwatch_rule.name
}*/
