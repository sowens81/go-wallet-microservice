resource "azurerm_cosmosdb_account" "cosmos" {
  name                = local.cosmos_name
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  offer_type          = "Standard"
  kind                = "GlobalDocumentDB"

  consistency_policy {
    consistency_level = "Session"
  }

  geo_location {
    location          = azurerm_resource_group.rg.location
    failover_priority = 0
  }
  dynamic "geo_location" {
    for_each = var.cosmos_geo.regions
    content {
      location          = geo_location.value
      failover_priority = index(var.cosmos_geo.regions, geo_location.value) + 1
    }
  }

  multiple_write_locations_enabled = var.cosmos_geo.multi_region_writes

  dynamic "capabilities" {
    for_each = var.cosmos_serverless ? [1] : []
    content {
      name = "EnableServerless"
    }
  }

  // backup configuration
  backup {
    type                = var.cosmos_backup.policy
    interval_in_minutes = var.cosmos_backup.interval_minutes
    retention_in_hours  = var.cosmos_backup.retention_hours
    storage_redundancy  = var.cosmos_backup.storage_redundancy
  }

  tags = var.tags
}

// create database if commit_id is null/empty
resource "azurerm_cosmosdb_sql_database" "db" {
  count               = var.commit_id == "" ? 1 : 0
  name                = var.database["name"]
  resource_group_name = azurerm_resource_group.rg.name
  account_name        = azurerm_cosmosdb_account.cosmos.name
  throughput          = var.cosmos_serverless ? null : var.database["throughput"]
}

// create container if commit_id is null/empty
resource "azurerm_cosmosdb_sql_container" "container" {
  count               = var.commit_id == "" ? 1 : 0
  name                = var.container["name"]
  resource_group_name = azurerm_resource_group.rg.name
  account_name        = azurerm_cosmosdb_account.cosmos.name
  database_name       = azurerm_cosmosdb_sql_database.db[0].name
  partition_key_paths = var.container["partition_key_paths"]
  throughput          = var.cosmos_serverless ? null : var.container["throughput"]
}




// create container if commit_id is null
