variable "replicationType" {
  description = "Type of replication"
  type        = string
  default     = "LRS"
}

variable "accountTier" {
  description = "Account Tier"
  type        = string
  default     = "Standard"
}

variable "accountKind" {
  description = "Account Kind"
  type        = string
  default     = "StorageV2"
}