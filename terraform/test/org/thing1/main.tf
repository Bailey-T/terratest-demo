resource "azurerm_resource_group" "rg" {
  name     = "terratest_${var.thing}"
  location = "uksouth"
}

resource "azurerm_private_dns_zone" "example" {
  name                = "${var.thing}.terratest.demo.local"
  resource_group_name = azurerm_resource_group.rg.name
}