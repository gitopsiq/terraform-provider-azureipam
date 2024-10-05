package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AZURE_IPAM_API_ENDPOINT", nil),
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

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	token := d.Get("token").(string)

	if token == "" {
			return nil, diag.Errorf("token is required")
	}

	client := NewClient(token, "https://api.azure-ipam.com")
	return client, diags
}

