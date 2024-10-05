package provider

import (
	"context"

	"github.com/gitopsiq/terraform-provider-azureipam/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns the schema.Provider for Azure IPAM
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AZURE_IPAM_API_ENDPOINT", "https://api.azure-ipam.com"),
				Description: "The Azure IPAM API endpoint URL",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AZURE_IPAM_TOKEN", nil),
				Description: "The Azure AD token for authentication",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azureipam_space": resourceSpace(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"azureipam_spaces": dataSourceSpaces(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// providerConfigure configures the API client used by the provider
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	apiEndpoint := d.Get("api_endpoint").(string)
	token := d.Get("token").(string)

	if token == "" {
		return nil, diag.Errorf("token is required")
	}

	if apiEndpoint == "" {
		return nil, diag.Errorf("API endpoint is required")
	}

	client := client.NewClient(token, apiEndpoint)

	return client, diags
}
