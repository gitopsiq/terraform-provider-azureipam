package provider

import (
    "context"
    "log"

    "github.com/gitopsiq/terraform-provider-azureipam/client"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "api_token": {
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("AZURE_IPAM_TOKEN", nil),
                Description: "The API token for authenticating with Azure IPAM",
            },
            "base_url": {
                Type:        schema.TypeString,
                Optional:    true,
                DefaultFunc: schema.EnvDefaultFunc("AZURE_IPAM_API_ENDPOINT", nil),
                Description: "The base URL for the Azure IPAM API",
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "azureipam_space": resourceSpace(),  // Register the space resource here
        },
        DataSourcesMap: map[string]*schema.Resource{
            "azureipam_spaces": dataSourceSpaces(),
        },
        ConfigureContextFunc: providerConfigure,
    }
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	baseURL := d.Get("base_url").(string)
	apiToken := d.Get("api_token").(string)  // Use the same 'api_token' as defined in schema

	// Initialize the client with the token
	c := client.NewClient(baseURL, apiToken)

	if c == nil {
			return nil, diag.Errorf("Unable to create Azure IPAM client")
	}

	log.Printf("[INFO] Azure IPAM Client configured with base URL: %s", baseURL)
	log.Printf("[DEBUG] Using API Token: %s", apiToken)  // Debug the token to ensure it's passed

	return c, diags
}


