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
    token := d.Get("api_token").(string)

    // Initialize the client
    c := client.NewClient(baseURL, token)

    if c == nil {
        return nil, diag.Errorf("Unable to create Azure IPAM client")
    }

    log.Printf("[INFO] Azure IPAM Client configured with base URL: %s", baseURL)

    return c, diags
}
