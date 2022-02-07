
resource "snowflake_table" "RAW_SOURCE" {
  database = snowflake_schema.SCHEMAJSON.database
  schema   = snowflake_schema.SCHEMAJSON.name
  name     = "RAW_SOURCE"
  column {
    name     = "SRC"
    type     = "VARIANT"
    nullable = false
  }

  column {
    name     = "V_SHA"
    type     = "VARCHAR(100)"
    nullable = false
  }
}


resource "snowflake_table" "USERS_TRANSFORMED" {
  database = snowflake_schema.SCHEMAJSON.database
  schema   = snowflake_schema.SCHEMAJSON.name
  name     = "USERS_TRANSFORMED"

  column {
    name     = "ID"
    type     = "VARCHAR(100)"
    nullable = false
  }

  column {
    name     = "FullName"
    type     = "VARCHAR(46)"
    nullable = false
  }

  column {
    name     = "Age"
    type     = "NUMBER(38,0)"
    nullable = false
  }

  column {
    name     = "City"
    type     = "VARCHAR(30)"
    nullable = false
  }

  column {
    name     = "Organization"
    type     = "VARCHAR(30)"
    nullable = false
  }

  column {
    name     = "Department"
    type     = "VARCHAR(30)"
    nullable = false
  }

  column {
    name     = "Subdepartment"
    type     = "VARCHAR(30)"
    nullable = false
  }

  column {
    name     = "Team"
    type     = "VARCHAR(30)"
    nullable = false
  }

}
