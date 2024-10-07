package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func resourceSpace() *schema.Resource {
	return &schema.Resource{
			CreateContext: resourceSpaceCreate,
			ReadContext:   resourceSpaceRead,
			DeleteContext: resourceSpaceDelete,
			Schema: map[string]*schema.Schema{
					"name": {
							Type:     schema.TypeString,
							Required: true,
					},
					"description": {
							Type:     schema.TypeString,
							Optional: true,
					},
			},
	}
}

func resourceSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	name := d.Get("name").(string)
	desc := d.Get("description").(string)

	space, err := c.CreateSpace(name, desc)
	if err != nil {
			return diag.FromErr(err)
	}

	d.SetId(space.ID)
	return resourceSpaceRead(ctx, d, m)
}