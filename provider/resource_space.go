package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/gitopsiq/terraform-provider-azureipam/client"
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
				Computed:    true, // Track optional description changes
				Description: "A description of the space",
			},
		},
	}
}

func resourceSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceName := d.Get("name").(string)
	spaceDesc := d.Get("description").(string)

	// Log the values being sent for creation
	log.Printf("[DEBUG] Creating space: name=%s, description=%s", spaceName, spaceDesc)

	// Create the space via API
	createdSpace, err := client.CreateSpace(spaceName, spaceDesc)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating space: %s", err))
	}

	// Set the resource ID
	d.SetId(createdSpace.ID)

	// Log the API response
	log.Printf("[DEBUG] API created space: ID=%s, name=%s, description=%s", createdSpace.ID, createdSpace.Name, createdSpace.Desc)

	// Set state values
	if err := d.Set("name", createdSpace.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", createdSpace.Desc); err != nil {  // Ensure description is set
		return diag.FromErr(err)
	}

	return resourceSpaceRead(ctx, d, m)
}

func resourceSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	// Log the Read operation
	log.Printf("[DEBUG] Reading space: ID=%s", spaceID)

	// Fetch space details via API
	space, err := client.GetSpace(spaceID)
	if err != nil {
		if space == nil {
			// Resource no longer exists
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	// Log the API response during read
	log.Printf("[DEBUG] API returned space: ID=%s, name=%s, description=%s", space.ID, space.Name, space.Desc)

	// Ensure Terraform state is consistent
	if err := d.Set("name", space.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", space.Desc); err != nil {  // Ensure description is set
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func resourceSpaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	if d.HasChange("description") {
		description := d.Get("description").(string)

		// Log the update
		log.Printf("[DEBUG] Updating space ID=%s with new description=%s", spaceID, description)

		// Update the space via API
		_, err := client.UpdateSpace(spaceID, description)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error updating space: %s", err))
		}
	}

	return resourceSpaceRead(ctx, d, m)
}

func resourceSpaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	// Log the deletion
	log.Printf("[DEBUG] Deleting space: ID=%s", spaceID)

	// Delete the space via API
	err := client.DeleteSpace(spaceID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting space: %s", err))
	}

	// Remove the resource from state
	d.SetId("")

	return diag.Diagnostics{}
}
