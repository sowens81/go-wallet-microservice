// Lookup the built-in role definition at subscription scope
resource "random_uuid" "role_assign_current" {}

data "azurerm_cosmosdb_sql_role_definition" "cosmos_db_data_contributor" {
  resource_group_name = azurerm_resource_group.rg.name
  account_name        = azurerm_cosmosdb_account.cosmos.name
  role_definition_id  = "00000000-0000-0000-0000-000000000002"
}

// Assign Cosmos SQL data role to the current principal (data plane)
resource "azurerm_cosmosdb_sql_role_assignment" "current_cosmos_data_writer" {
  name                = random_uuid.role_assign_current.result
  account_name        = azurerm_cosmosdb_account.cosmos.name
  scope               = azurerm_cosmosdb_account.cosmos.id
  resource_group_name = azurerm_resource_group.rg.name
  principal_id        = data.azurerm_client_config.current.object_id
  role_definition_id  = data.azurerm_cosmosdb_sql_role_definition.cosmos_db_data_contributor.id
  depends_on = [
    azurerm_cosmosdb_account.cosmos
  ]
}

// Assign Cosmos SQL data role to the managed identity
resource "random_uuid" "role_assign_mi" {}

resource "azurerm_cosmosdb_sql_role_assignment" "mi_cosmos_data_writer" {
  name                = random_uuid.role_assign_mi.result
  account_name        = azurerm_cosmosdb_account.cosmos.name
  scope               = azurerm_cosmosdb_account.cosmos.id
  resource_group_name = azurerm_resource_group.rg.name
  principal_id        = azurerm_user_assigned_identity.mi.principal_id
  role_definition_id  = data.azurerm_cosmosdb_sql_role_definition.cosmos_db_data_contributor.id
  depends_on = [
    azurerm_cosmosdb_account.cosmos,
    azurerm_user_assigned_identity.mi
  ]
}
