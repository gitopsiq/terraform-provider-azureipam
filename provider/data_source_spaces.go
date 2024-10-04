package provider

import (
	"context"
	"strconv"
	"time"

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
	// client := m.(*Client)

	// TODO: Use the client to fetch spaces
	// spaces, err := client.ListSpaces(ctx)
	// if err != nil {
	//     return diag.FromErr(err)
	// }

	// For now, we'll use mock data
	spaces := []map[string]interface{}{
		{
			"name":        "space1",
			"description": "This is space 1",
		},
		{
			"name":        "space2",
			"description": "This is space 2",
		},
	}

	if err := d.Set("spaces", spaces); err != nil {
		return diag.FromErr(err)
	}

	// Generate a unique ID for this data source
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diag.Diagnostics{}
}