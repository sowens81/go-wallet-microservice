output "resource_group_name" {
  description = "Resource group name"
  value       = azurerm_resource_group.rg.name
}

output "resource_group_id" {
  description = "Resource group id"
  value       = azurerm_resource_group.rg.id
}

output "managed_identity_name" {
  description = "Managed identity name"
  value       = azurerm_user_assigned_identity.mi.name
}

output "managed_identity_id" {
  description = "Managed identity client id"
  value       = azurerm_user_assigned_identity.mi.client_id
}

output "cosmos_endpoint" {
  description = "Cosmos DB Endpoint"
  value       = azurerm_cosmosdb_account.cosmos.endpoint
}

output "cosmos_database_name" {
  description = "Cosmos DB Database Name"
  value       = length(azurerm_cosmosdb_sql_database.db) > 0 ? azurerm_cosmosdb_sql_database.db[0].name : null
}

output "cosmos_container_name" {
  description = "Cosmos DB Container Name"
  value       = length(azurerm_cosmosdb_sql_container.container) > 0 ? azurerm_cosmosdb_sql_container.container[0].name : null
}
