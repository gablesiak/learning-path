resource "snowflake_task" "copy_task" {

  database  = snowflake_schema.SCHEMAJSON.database
  schema    = snowflake_schema.SCHEMAJSON.name
  warehouse = "TEST_WH"

  name          = "copy_task"
  schedule      = "10 MINUTE"
  sql_statement = "copy into RAW_SOURCE(SRC, V_SHA) from (select $1, SHA2($1) from @azure_stage_sas)"
}


resource "snowflake_task" "insert_data" {

  database  = snowflake_schema.SCHEMAJSON.database
  schema    = snowflake_schema.SCHEMAJSON.name
  warehouse = "TEST_WH"

  name          = "insert_data"
  schedule      = "10 MINUTE"
  sql_statement = "insert into USERS_TRANSFORMED select V_SHA, src:FULLNAME, src:AGE, src:CITY, src:ORGANIZATION, src:DEPARTMENT, src:SUBDEPARTMENT, src:TEAM from RAW_SOURCE where not exists (select 'x' from USERS_TRANSFORMED where USERS_TRANSFORMED.ID = RAW_SOURCE.V_SHA);"
}