<#
.SYNOPSIS
  Deploy Terraform and generate a .env file with Cosmos outputs.

.DESCRIPTION
  This script runs `terraform init`, `terraform plan -var-file=env.tfvars -out env.tfplan`,
  and `terraform apply env.tfplan` in the `terraform/` directory. It then reads the
  Terraform outputs `cosmos_endpoint`, `cosmos_database_name` and `cosmos_container_name`
  and writes a `.env` file at the repository root with the keys:
    COSMOS_ENDPOINT, COSMOS_DB, COSMOS_CONTAINER

.PARAMETER VarFile
  Optional path to a tfvars file. Defaults to `terraform/env.tfvars`.

.EXAMPLE
  .\scripts\deploy.ps1
  .\scripts\deploy.ps1 -VarFile .\terraform\env.tfvars
#>

[CmdletBinding()]
param(
    [string]$VarFile = "$PSScriptRoot/../terraform/env.tfvars"
)

function FailIf([int]$code, [string]$msg) {
    if ($code -ne 0) { Write-Error $msg; exit $code }
}

$RepoRoot = Resolve-Path "$PSScriptRoot/.."
$TfDir = Join-Path $RepoRoot "terraform"

if (-not (Get-Command terraform -ErrorAction SilentlyContinue)) {
    Write-Error "terraform not found in PATH"
    exit 1
}

Write-Host "Using tfvars: $VarFile"
Push-Location $TfDir

Write-Host "terraform init"
terraform init -input=false
FailIf $LASTEXITCODE "terraform init failed"

Write-Host "terraform plan -var-file=$VarFile -out env.tfplan"
& terraform plan -var-file "$VarFile" -out "env.tfplan"
FailIf $LASTEXITCODE "terraform plan failed"

Write-Host "terraform apply env.tfplan"
& terraform apply -input=false "env.tfplan"
FailIf $LASTEXITCODE "terraform apply failed"

Write-Host "Reading Terraform outputs"
$tfJson = terraform output -json | ConvertFrom-Json

$endpoint = if ($tfJson.cosmos_endpoint -ne $null) { $tfJson.cosmos_endpoint.value } else { "" }
$dbname   = if ($tfJson.cosmos_database_name -ne $null) { $tfJson.cosmos_database_name.value } else { "" }
$container= if ($tfJson.cosmos_container_name -ne $null) { $tfJson.cosmos_container_name.value } else { "" }

$envPath = Join-Path $RepoRoot ".env"
Set-Content -Path $envPath -Value @(
    "COSMOS_ENDPOINT=$endpoint",
    "COSMOS_DB=$dbname",
    "COSMOS_CONTAINER=$container"
)

Write-Host "Wrote $envPath"
Pop-Location

Exit 0
