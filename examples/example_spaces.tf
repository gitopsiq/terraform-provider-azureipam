provider "azureipam" {}

data "azureipam_spaces" "all" {}

output "spaces" {
  value = data.azureipam_spaces.all.spaces
}