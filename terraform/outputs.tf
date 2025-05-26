output "resource_group_name" {
  value = azurerm_resource_group.pso_rg.name
}

output "public_ip_address" {
  value = azurerm_linux_virtual_machine.pso_vm.public_ip_address
}