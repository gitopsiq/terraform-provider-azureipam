package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/gitopsiq/terraform-provider-azureipam/client"
)

// resourceSpace returns the schema for the IPAM Space resource.
func resourceSpace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSpaceCreate,
		ReadContext:   resourceSpaceRead,
		UpdateContext: resourceSpaceUpdate,
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

// resourceSpaceCreate creates a new space in Azure IPAM.
func resourceSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	name := d.Get("name").(string)
	desc := d.Get("description").(string)

	// Call client to create the space
	space, err := c.CreateSpace(ctx, name, desc)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the resource ID (space ID) and sync data with Terraform
	d.SetId(space.ID)

	// Call read to sync state with the created resource
	return resourceSpaceRead(ctx, d, m)
}

// resourceSpaceRead reads the current state of the space from Azure IPAM.
func resourceSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	spaceID := d.Id()

	// Fetch the space from Azure IPAM using the space ID
	space, err := c.GetSpace(ctx, spaceID)
	if err != nil {
		// If the space is not found, remove from state
		if err.Error() == "not found" {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	// Sync the state with Terraform
	d.Set("name", space.Name)
	d.Set("description", space.Description)

	return diag.Diagnostics{}
}

// resourceSpaceUpdate updates the existing space in Azure IPAM.
func resourceSpaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	spaceID := d.Id()

	// Only update if there are changes
	if d.HasChange("name") || d.HasChange("description") {
		name := d.Get("name").(string)
		desc := d.Get("description").(string)

		// Call client to update the space
		_, err := c.UpdateSpace(ctx, spaceID, name, desc)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Call read to sync state after the update
	return resourceSpaceRead(ctx, d, m)
}

// resourceSpaceDelete deletes a space in Azure IPAM.
func resourceSpaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	spaceID := d.Id()

	// Call client to delete the space
	err := c.DeleteSpace(ctx, spaceID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Remove from Terraform state
	d.SetId("")

	return diag.Diagnostics{}
}
