package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"azureipam": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example:
	if v := os.Getenv("AZURE_IPAM_API_ENDPOINT"); v == "" {
		t.Fatal("AZURE_IPAM_API_ENDPOINT must be set for acceptance tests")
	}
	if v := os.Getenv("AZURE_IPAM_TOKEN"); v == "" {
		t.Fatal("AZURE_IPAM_TOKEN must be set for acceptance tests")
	}
}