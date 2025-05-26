resource "random_id" "random_id" {
  keepers = {
    # Generate a new ID only when a new resource group is defined
    resource_group = azurerm_resource_group.pso_rg.name
  }

  byte_length = 8
}

resource "azurerm_storage_account" "pso_sg" {
  name                     = "diag${random_id.random_id.hex}"
  location                 = azurerm_resource_group.pso_rg.location
  resource_group_name      = azurerm_resource_group.pso_rg.name
  account_tier             = "Standard"
  account_replication_type = "LRS"
}