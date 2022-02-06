terraform {
  required_providers {
    snowflake = {
      source  = "chanzuckerberg/snowflake"
      version = "0.25.36"
    }
  }
}

provider "snowflake" {
  username = "${var.snowflake_username}"
  account  = "${var.snowflake_account}"
  role     = "ACCOUNTADMIN"
  password = "${var.snowflake_password}"
}

resource "snowflake_warehouse" "TEST_WH" {
  name           = "TEST_WH"
  warehouse_size = "xsmall"
}

resource "snowflake_database" "BLOB_DB" {
  name                        = "BLOB_DB"
  data_retention_time_in_days = 1
}

resource "snowflake_schema" "SCHEMAJSON" {
  name     = "SCHEMAJSON"
  database = "BLOB_DB"
}


