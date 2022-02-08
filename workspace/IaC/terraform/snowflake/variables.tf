variable "snowflake_database" {
  default="BLOB_DB"
}

variable "snowflake_schema"{
  default="PUBLIC"
}

variable "az_sas_token" {
  type = string
}

variable "snowflake_username" {
  type = string
}

variable "snowflake_account" {
  type = string
}

variable "snowflake_password" {
  type = string
}

variable "az_container_url" {
  type = string
}