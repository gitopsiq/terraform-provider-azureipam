
# Azure IPAM Terraform Provider

This is a Terraform provider for Azure IP Address Management (IPAM). It allows you to manage IP address spaces, blocks, and reservations using Terraform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.9.x or later
- [Go](https://golang.org/doc/install) 1.23 or later (for development)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) (for obtaining authentication token)
- [Azure PowerShell](https://docs.microsoft.com/en-us/powershell/azure/new-azureps-module-az)

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```sh
go install
```

## Using the provider

### Authentication

To use this provider, you need to authenticate with Azure IPAM and cache the authentication token and API endpoint for future use.

### Step 1: PowerShell Script for Authentication, App Service Discovery, and Token Caching

Run the following **PowerShell** script to authenticate, discover the `Application ID` and App Service URL, and generate an Azure AD token. The script will save the token and API endpoint in `~/.ipam/.token`.

```powershell
# Step 1: Authenticate with Azure
Connect-AzAccount

# Step 2: Discover the Application ID for the ipam-engine-app
$engineAppName = "ipam-engine-app"
$engineApp = Get-AzADApplication -Filter "DisplayName eq '$engineAppName'"
$engineAppClientId = $engineApp.AppId

# Step 3: Generate the Azure AD token
$accessToken = (Get-AzAccessToken -ResourceUrl "api://$engineAppClientId").Token

# Step 4: Discover the App Service URL that starts with 'ipam'
$appServiceUrl = az webapp list --query "[?starts_with(name, 'ipam')].defaultHostName" -o tsv

# Step 5: Create a folder ~/.ipam if it doesn't exist
$ipamFolderPath = "$HOME/.ipam"
if (!(Test-Path -Path $ipamFolderPath)) {
    New-Item -Path $ipamFolderPath -ItemType Directory
}

# Step 6: Write the token and API endpoint to ~/.ipam/.token
$tokenFilePath = "$ipamFolderPath/.token"
$tokenContent = @"
AZURE_IPAM_TOKEN=$accessToken
AZURE_IPAM_API_ENDPOINT=https://$appServiceUrl
"@
Set-Content -Path $tokenFilePath -Value $tokenContent

# Confirmation message
Write-Host "Token and API endpoint saved to $tokenFilePath"
```

### Step 2: Bash Script for Loading Cached Token and API Endpoint

After running the PowerShell script, you can load the token and API endpoint from the cached file in `~/.ipam/.token` in **Bash**.

```bash
# Load the token and API endpoint from the cache
if [ -f ~/.ipam/.token ]; then
    source ~/.ipam/.token
    echo "Token and API endpoint loaded from ~/.ipam/.token"
else
    echo "Token file not found. Please run the PowerShell script to authenticate and cache the token."
fi

# Verify that the environment variables are set
echo "Bash environment variables for provider authentications set:"
env | grep AZURE_IPAM
```

### Step 3: Use the Provider in Terraform

Once the token and API endpoint are loaded, you can use the provider in your Terraform configuration:

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
  api_endpoint = "${env.AZURE_IPAM_API_ENDPOINT}"
  token        = "${env.AZURE_IPAM_TOKEN}"
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
export AZURE_IPAM_API_ENDPOINT="https://your-ipam-instance.azurewebsites.net"
export AZURE_IPAM_TOKEN="your-azure-ad-token"
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---
