package provider

import (
    "context"

    "github.com/gitopsiq/terraform-provider-azureipam/client"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSpaces() *schema.Resource {
    return &schema.Resource{
        ReadContext: dataSourceSpacesRead,
        Schema: map[string]*schema.Schema{
            "spaces": {
                Type:     schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type:     schema.TypeString,
                            Computed: true,
                        },
                        "name": {
                            Type:     schema.TypeString,
                            Computed: true,
                        },
                        "description": {
                            Type:     schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

// dataSourceSpacesRead reads and returns the list of spaces from Azure IPAM.
func dataSourceSpacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    c := m.(*client.Client)

    var diags diag.Diagnostics

    // Fetch spaces using the context
    spaces, err := c.GetAllSpaces(ctx)
    if err != nil {
        return diag.FromErr(err)
    }

    // Flatten the spaces data and set it in the Terraform state
    if err := d.Set("spaces", flattenSpacesData(spaces)); err != nil {
        return diag.FromErr(err)
    }

    // Set a fixed ID to ensure Terraform detects changes properly
    d.SetId("spaces-list")

    return diags
}

// Helper function to flatten spaces data into a list for Terraform
func flattenSpacesData(spaces []client.Space) []interface{} {
    if spaces != nil {
        results := make([]interface{}, len(spaces))

        for i, space := range spaces {
            result := make(map[string]interface{})
            result["id"] = space.ID
            result["name"] = space.Name
            result["description"] = space.Description

            results[i] = result
        }

        return results
    }

    return make([]interface{}, 0)
}
