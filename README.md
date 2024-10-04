# Azure IPAM Terraform Provider

This is a Terraform provider for Azure IP Address Management (IPAM). It allows you to manage IP address spaces, blocks, and reservations using Terraform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or later
- [Go](https://golang.org/doc/install) 1.23 or later (for development)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) (for obtaining authentication token)

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```sh
go install
```

## Using the provider

### Authentication

To use this provider, you need to authenticate with Azure. You can obtain an Azure AD token using the Azure CLI. Here's how:

1. Ensure you have the Azure CLI installed and you're logged in:

   ```sh
   az login
   ```

2. If you have multiple subscriptions, select the one you want to use:

   ```sh
   az account set --subscription="SUBSCRIPTION_ID"
   ```

3. Obtain an access token:

   ```sh
   export AZURE_IPAM_TOKEN=$(az account get-access-token --resource="https://management.azure.com/" --query accessToken -o tsv)
   ```

   Replace `https://management.azure.com/` with the appropriate resource URL if your IPAM API uses a different resource.

4. Set the API endpoint:

   ```sh
   export AZURE_IPAM_API_ENDPOINT="https://your-api-endpoint.com"
   ```

Now you can use the provider in your Terraform configuration:

```hcl
terraform {
  required_providers {
    azureipam = {
      source = "gitopsiq/azureipam"
    }
  }
}

provider "azureipam" {}
```

Alternatively, you can provide the token and API endpoint directly in the provider configuration:

```hcl
provider "azureipam" {
  api_endpoint = "https://your-api-endpoint.com"
  token        = "your-azure-ad-token"
}
```

### Available Resources

- `azureipam_space`: Manages an IPAM space

Example usage:

```hcl
resource "azureipam_space" "example" {
  name        = "example-space"
  description = "This is an example space"
}
```

### Available Data Sources

- `azureipam_spaces`: Retrieves a list of IPAM spaces

Example usage:

```hcl
data "azureipam_spaces" "all" {}

output "all_spaces" {
  value = data.azureipam_spaces.all.spaces
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.23+ is *required*).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ go install
```

### Testing the Provider

To run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

To run the tests without `make`, you can use:

```sh
$ TF_ACC=1 go test ./... -v
```

Make sure to set the required environment variables before running the tests:

```sh
export AZURE_IPAM_API_ENDPOINT="https://your-api-endpoint.com"
export AZURE_IPAM_TOKEN="your-azure-ad-token"
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.