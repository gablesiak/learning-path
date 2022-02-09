
resource "snowflake_table" "RAW_SOURCE" {
  database = var.snowflake_database
  schema   = var.snowflake_schema
  name     = "RAW_SOURCE"
  column {
    name     = "SRC"
    type     = "VARIANT"
    nullable = false
  }

  column {
    name     = "SHA"
    type     = "VARCHAR(100)"
    nullable = false
  }
}


resource "snowflake_table" "USERS_TRANSFORMED" {
  database = var.snowflake_database
  schema   = var.snowflake_schema
  name     = "USERS_TRANSFORMED"

  column {
    name     = "SHA"
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
