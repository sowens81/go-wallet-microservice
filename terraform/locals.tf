locals {
  // short mappings for environment
  env_short_map = {
    production  = "prd"
    staging     = "stg"
    development = "dev"
  }

  // short mappings for regions (extend as needed)
  region_short_map = {
    uksouth     = "uks"
    ukwest      = "ukw"
    westeurope  = "weu"
    northeurope = "neu"
  }

  envShort    = lookup(local.env_short_map, lower(var.environment), substr(lower(var.environment), 0, 3))
  regionShort = lookup(local.region_short_map, lower(var.region), substr(lower(var.region), 0, 3))

  commit_suffix = var.commit_id != "" ? format("-%s", var.commit_id) : ""

  // resource name helper: {org}-{envShort}-{regionShort}-rg-{resource_name}{-commit}
  rg_name     = format("%s-%s-%s-rg-%s%s", var.organisation, local.envShort, local.regionShort, var.service_name, local.commit_suffix)
  cosmos_name = format("%s-%s-%s-rg-%s%s", var.organisation, local.envShort, local.regionShort, var.service_name, local.commit_suffix)
  mi_name     = format("%s-%s-%s-rg-%s%s", var.organisation, local.envShort, local.regionShort, var.service_name, local.commit_suffix)
}
