package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/gitopsiq/terraform-provider-azureipam/client" // Ensure you import the client package
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
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the space",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the space",
						},
					},
				},
			},
		},
	}
}

func dataSourceSpacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// Retrieve the client from the provider configuration
	client := m.(*client.Client)

	// Fetch the list of spaces from the API
	spaces, err := client.ListSpaces()
	if err != nil {
		return diag.FromErr(err)
	}

	// Convert the spaces into a format suitable for Terraform
	var spaceList []map[string]interface{}
	for _, space := range spaces {
		spaceData := map[string]interface{}{
			"name":        space.Name,
			"description": space.Desc, // Assuming the API uses 'Desc' for description
		}
		spaceList = append(spaceList, spaceData)
	}

	// Set the spaces in the Terraform state
	if err := d.Set("spaces", spaceList); err != nil {
		return diag.FromErr(err)
	}

	// Set a unique ID for this data source, based on the current timestamp
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
