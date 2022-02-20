 # Connecting to azure

Connect-AzureRmAccount

# Setting env variables from keyvault



$env:PULUMI_snowflake_username = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "snowflakeusername").secretvaluetext
$env:PULUMI_snowflake_account = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "snowflakeaccount").secretvaluetext
$env:PULUMI_snowflake_password = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "snowflakepassword").secretvaluetext
$env:PULUMI_az_sas_token = (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "az-sastoken").secretvaluetext
$env:PULUMI_az_container_url= (Get-AzureKeyVaultSecret -VaultName "snowflakedata" -Name "azcontainerurl").secretvaluetext

pulumi config set snowflake:username $env:PULUMI_snowflake_username
pulumi config set snowflake:account $env:PULUMI_snowflake_account
pulumi config set snowflake:password $env:PULUMI_snowflake_password --secret
pulumi config set azure-native:az_sas_token $env:PULUMI_az_sas_token --secret
pulumi config set azure-native:az_container_url $env:PULUMI_az_container_url

# Initializing pulumi

pulumi preview

pulumi up --yes

# Removing env variables

Get-ChildItem -Path Env:\ | Where-Object {$_.Name -like 'PULUMI_*'} | Remove-Item
