provider "azurerm" {
    features {}
}

data "azurerm_resource_group" "ResourceGroup" {
  name = "cloud-shell-storage-westeurope"
}

resource "azurerm_storage_account" "StorageAccount" {
  name                     = "learning"
  resource_group_name      = data.azurerm_resource_group.ResourceGroup.name
  location                 = data.azurerm_resource_group.ResourceGroup.location
  account_kind             = var.accountKind
  account_tier             = var.accountTier
  account_replication_type = var.replicationType
}

resource "azurerm_storage_container" "Container" {
  name                  = "rel-project"
  storage_account_name  = azurerm_storage_account.StorageAccount.name
  container_access_type = "private"
}
