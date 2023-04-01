resource "azurerm_resource_group" "rg" {
  name     = "terratest_${var.thing}_${var.guid}"
  location = "uksouth"
}

resource "azurerm_private_dns_zone" "example" {
  name                = "${var.thing}.${var.guid}.terratest.demo.local"
  resource_group_name = azurerm_resource_group.rg.name
}