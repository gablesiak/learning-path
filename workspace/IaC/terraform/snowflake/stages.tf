resource "snowflake_stage" "azure_stage_sas" {
  name        = "azure_stage_sas"
  url         = "${var.az_container_url}"
  database    = snowflake_schema.SCHEMAJSON.database
  schema      = snowflake_schema.SCHEMAJSON.name
  credentials = "AZURE_SAS_TOKEN='${var.az_sas_token}'"
}