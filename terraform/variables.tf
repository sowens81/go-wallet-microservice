variable "organisation" {
  type        = string
  description = "Organisation name used in resource naming"
}

variable "environment" {
  type        = string
  description = "Deployment environment (e.g. production, staging, development)"
}

variable "region" {
  type        = string
  description = "Azure region (e.g. uksouth, eastus)"
  default     = "uksouth"
}

variable "service_name" {
  type        = string
  description = "Service name used in resource naming"
}

variable "commit_id" {
  type        = string
  description = "Optional commit id to append to resource names"
  default     = ""
}

variable "tags" {
  type        = map(string)
  description = "Tags to apply to all resources"
  default     = {}
}

variable "cosmos_serverless" {
  type        = bool
  description = "Enable Cosmos DB serverless mode (default true)"
  default     = true
}

variable "cosmos_geo" {
  type = object({
    geo_redundancy_enabled = bool
    multi_region_writes    = bool
    regions                = list(string)
  })
  description = "Geo / replication settings for Cosmos DB (regions list is empty by default)"
  default = {
    geo_redundancy_enabled = false
    multi_region_writes    = false
    regions                = []
  }
}

variable "cosmos_backup" {
  type = object({
    policy             = string
    interval_minutes   = number
    retention_hours    = number
    storage_redundancy = string
  })
  description = "Backup configuration for Cosmos DB"
  default = {
    policy             = "Periodic"
    interval_minutes   = 1440
    retention_hours    = 48
    storage_redundancy = "Local"
  }
}

variable "database" {
  type = object({
    name       = string
    throughput = number
  })
  description = "Cosmos DB SQL database settings"
  default = {
    name       = ""
    throughput = 400
  }
}

variable "container" {
  type = object({
    name                = string
    partition_key_paths = list(string)
    throughput          = number
  })
  description = "Cosmos DB SQL container settings"
  default = {
    name                = ""
    partition_key_paths = ["/id"]
    throughput          = 400
  }
}


