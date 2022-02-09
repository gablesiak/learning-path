resource "snowflake_task" "copy_data_from_azure" {

  database = var.snowflake_database
  schema   = var.snowflake_schema
  warehouse = "TEST_WH"

  name          = "copy_data_from_azure"
  sql_statement = "copy into RAW_SOURCE(SRC, SHA) from (select $1, SHA2($1) from @azure_stage_sas)"
}


resource "snowflake_task" "insert_data_into_users_transformed" {

  database = var.snowflake_database
  schema   = var.snowflake_schema
  warehouse = "TEST_WH"

  name          = "insert_data_into_users_transformed"
  schedule      = "10 MINUTE"
  sql_statement = "insert into USERS_TRANSFORMED select SHA, src:FULLNAME, src:AGE, src:CITY, src:ORGANIZATION, src:DEPARTMENT, src:SUBDEPARTMENT, src:TEAM from RAW_SOURCE where not exists (select 'x' from USERS_TRANSFORMED where USERS_TRANSFORMED.SHA = RAW_SOURCE.SHA);"
}