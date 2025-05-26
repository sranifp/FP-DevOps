resource "random_pet" "pso_rg_name" {
  prefix = var.resource_group_name_prefix
}

resource "azurerm_resource_group" "pso_rg" {
  location = var.resource_group_location
  name     = random_pet.pso_rg_name.id
}