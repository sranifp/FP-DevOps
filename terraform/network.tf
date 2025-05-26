resource "azurerm_virtual_network" "pso_vnet" {
  name                = "pso_vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.pso_rg.location
  resource_group_name = azurerm_resource_group.pso_rg.name
}

resource "azurerm_subnet" "pso_subnet" {
  name                 = "pso_subnet"
  resource_group_name  = azurerm_resource_group.pso_rg.name
  virtual_network_name = azurerm_virtual_network.pso_vnet.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_public_ip" "pso_public_ip" {
  name                = "pso_public_ip"
  location            = azurerm_resource_group.pso_rg.location
  resource_group_name = azurerm_resource_group.pso_rg.name
  allocation_method   = "Dynamic"
}

resource "azurerm_network_security_group" "pso_nsg" {
  name                = "pso_nsg"
  location            = azurerm_resource_group.pso_rg.location
  resource_group_name = azurerm_resource_group.pso_rg.name

  security_rule {
    name                       = "SSH"
    priority                   = 1001
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "22"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "HTTP"
    priority                   = 1001
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "80"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}

resource "azurerm_network_interface" "pso_nic" {
  name                = "pso_nic"
  location            = azurerm_resource_group.pso_rg.location
  resource_group_name = azurerm_resource_group.pso_rg.name

  ip_configuration {
    name                          = "pso_nic_config"
    subnet_id                     = azurerm_subnet.pso_subnet.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.pso_public_ip.id
  }
}

resource "azurerm_network_interface_security_group_association" "pso_secgroup" {
  network_interface_id      = azurerm_network_interface.pso_nic.id
  network_security_group_id = azurerm_network_security_group.pso_nsg.id
}