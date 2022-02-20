# Connecting to azure

az login 

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
