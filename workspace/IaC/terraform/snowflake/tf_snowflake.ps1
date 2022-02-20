# Connecting to azure

Connect-AzureRmAccount

# Setting env variables from keyvault

$env:TF_VAR_snowflake_username = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "snowflakeusername").secretvaluetext
$env:TF_VAR_snowflake_account = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "snowflakeaccount").secretvaluetext
$env:TF_VAR_snowflake_password = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "snowflakepassword").secretvaluetext
$env:TF_VAR_az_sas_token = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "az-sastoken").secretvaluetext
$env:TF_VAR_az_container_url= (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "azcontainerurl").secretvaluetext

# Initializing terraform

terraform init

terraform plan

try {
terraform apply -auto-approve
}
finally {
  if(Test-Path "terraform.tfstate"){
   Remove-Item -Path "terraform.tfstate"
  }

  if(Test-Path "terraform.tfstate.backup"){
   Remove-Item -Path "terraform.tfstate.backup"
  }

  if(Test-Path ".terraform.lock.hcl"){
   Remove-Item -Path ".terraform.lock.hcl"
  }

  if(Test-Path ".terraform/"){
  Remove-Item -Path ".terraform" -Recurse
  }
}


# Removing env variables

Get-ChildItem -Path Env:\ | Where-Object {$_.Name -like 'TF_VAR_*'} | Remove-Item
