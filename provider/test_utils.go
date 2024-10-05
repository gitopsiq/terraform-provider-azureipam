package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// testAccPreCheck ensures that necessary environment variables are set before running acceptance tests
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("AZURE_IPAM_TOKEN"); v == "" {
		t.Fatal("AZURE_IPAM_TOKEN must be set for acceptance tests")
	}
}

// testAccProviders returns a map of providers for testing
func testAccProviders(apiURL string) map[string]*schema.Provider {
	return map[string]*schema.Provider{
		"azureipam": ProviderWithApiEndpoint(apiURL),
	}
}

// ProviderWithApiEndpoint configures the provider with a custom API endpoint (used for testing)
func ProviderWithApiEndpoint(apiEndpoint string) *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: func() (interface{}, error) {
					return apiEndpoint, nil
				},
				Description: "The Azure IPAM API endpoint URL",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The Azure AD token for authentication",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azureipam_space": resourceSpace(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"azureipam_spaces": dataSourceSpaces(),
		},
	}
}
