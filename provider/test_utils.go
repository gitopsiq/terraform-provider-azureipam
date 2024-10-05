package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// testAccProviders returns a map of providers for testing
func testAccProviders(apiURL string) map[string]*schema.Provider {
	return map[string]*schema.Provider{
		"azureipam": ProviderWithApiEndpoint(apiURL),
	}
}
