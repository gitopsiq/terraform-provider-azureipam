package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSpace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSpaceCreate,
		ReadContext:   resourceSpaceRead,
		UpdateContext: resourceSpaceUpdate,
		DeleteContext: resourceSpaceDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the space",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the space",
			},
		},
	}
}

func resourceSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	// TODO: Use the client to create the space
	// For now, we'll just use the name and description to show they're being used
	d.SetId(name)
	d.Set("description", description)

	return resourceSpaceRead(ctx, d, m)
}

func resourceSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: Implement the read operation
	// This would typically involve calling the API to get the current state of the resource
	// and updating the Terraform state with the returned data

	// For now, we'll just return an empty diagnostics slice
	return diag.Diagnostics{}
}

func resourceSpaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Check if the description has changed
	if d.HasChange("description") {
		description := d.Get("description").(string)
		// TODO: Use the client to update the space's description
		// For now, we'll just update the local state
		d.Set("description", description)
	}

	return resourceSpaceRead(ctx, d, m)
}

func resourceSpaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: Implement the delete operation
	// This would typically involve calling the API to delete the resource

	// For now, we'll just clear the ID to remove it from the Terraform state
	d.SetId("")

	return diag.Diagnostics{}
}