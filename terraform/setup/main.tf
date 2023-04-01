resource "azurerm_resource_group" "managed_identity" {
  name     = "managed_identity"
  location = "uksouth"
}

resource "azurerm_user_assigned_identity" "github" {
  name                = "tom-b-github-actions-workload-identity"
  resource_group_name = azurerm_resource_group.managed_identity.name
  location            = azurerm_resource_group.managed_identity.location
}

resource "azurerm_role_assignment" "github_actions_mi" {
  scope                = "/subscriptions/${data.azurerm_client_config.local.subscription_id}/"
  role_definition_name = "Contributor"
  principal_id         = azurerm_user_assigned_identity.github.principal_id
}

resource "azurerm_federated_identity_credential" "github_workload_identity" {
  name                = "${azurerm_user_assigned_identity.github.name}-workload-identity"
  resource_group_name = azurerm_resource_group.managed_identity.name
  audience            = ["api://AzureADTokenExchange"]
  issuer              = "https://token.actions.githubusercontent.com"
  parent_id           = azurerm_user_assigned_identity.github.id
  subject             = "repo:Bailey-T/terratest-demo:environment:azure-lab"
}