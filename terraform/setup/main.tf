resource "azurerm_resource_group" "managed_identity" {
  name = "managed_identity"
  location = "uksouth"  
}

resource "azurerm_user_assigned_identity" "github" {
  name = "tom-b-github-actions-workload-identity"
  resource_group_name = azurerm_resource_group.managed_identity.name
  location = azurerm_resource_group.managed_identity.location
}

resource "azurerm_federated_identity_credential" "github_workload_identity" {
  name = "${azurerm_user_assigned_identity.github.name}-workload-identity"
  resource_group_name = azurerm_resource_group.managed_identity.name
  audience = ["123"]
  issuer = "abc" 
  parent_id = azurerm_resource_group.managed_identity.id
  subject = "def"
}