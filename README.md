# WORK IN PROGRESS - Azure IPAM Terraform Provider

This is a Terraform provider for Azure IP Address Management (IPAM). It allows you to manage IP address spaces, blocks, and reservations using Terraform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or later
- [Go](https://golang.org/doc/install) 1.20 or later (for development)

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```sh
go install
```

## Using the provider

To use the provider, add the following to your Terraform configuration:

```hcl
terraform {
  required_providers {
    azureipam = {
      source = "gitopsiq/azureipam"
    }
  }
}

provider "azureipam" {
  api_endpoint = "https://your-api-endpoint.com"
  token        = "your-azure-ad-token"
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.20+ is *required*).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ go install
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
