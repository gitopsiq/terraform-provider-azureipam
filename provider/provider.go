package provider

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "github.com/gitopsiq/terraform-provider-azureipam/client"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
            "azureipam_space": resourceSpace(),
        },
        DataSourcesMap: map[string]*schema.Resource{
            "azureipam_spaces": dataSourceSpaces(),
        },
        ConfigureContextFunc: providerConfigure,
    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, diag.Diagnostics) {
    baseURL := d.Get("base_url").(string)
    token := d.Get("token").(string)

    c := client.NewClient(baseURL, token)
    return c, nil
}