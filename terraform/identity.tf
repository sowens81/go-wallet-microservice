// data source for the current principal (whoever runs terraform)
data "azurerm_client_config" "current" {}

resource "azurerm_user_assigned_identity" "mi" {
  name                = local.mi_name
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  tags                = var.tags
}