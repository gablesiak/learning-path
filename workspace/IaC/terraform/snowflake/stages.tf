resource "snowflake_stage" "azure_stage_sas" {
  name        = "azure_stage_sas"
  url         = var.az_container_url
  database = var.snowflake_database
  schema   = var.snowflake_schema
  credentials = "AZURE_SAS_TOKEN='${var.az_sas_token}'"
}