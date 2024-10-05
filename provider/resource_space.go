package provider

import (
	"context"
	"fmt"  // Add fmt to handle error formatting
	"github.com/gitopsiq/terraform-provider-azureipam/client" 
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

// Fix the function signature to include context
func resourceSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceName := d.Get("name").(string)
	spaceDesc := d.Get("description").(string)  // Use "description" for consistency

	// Call the API to create a space
	createdSpace, err := client.CreateSpace(spaceName, spaceDesc)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating space: %w", err))
	}

	// Set the ID for the resource
	d.SetId(createdSpace.ID)

	// Call resourceSpaceRead to sync the state after creation
	return resourceSpaceRead(ctx, d, m)
}

func resourceSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	// Call the API to get the space details
	space, err := client.GetSpace(spaceID)
	if err != nil {
		// Handle the case where the space was not found
		if space == nil {
			d.SetId("")
			return diag.FromErr(fmt.Errorf("space not found: %w", err))
		}
		return diag.FromErr(fmt.Errorf("error reading space: %w", err))
	}

	// Update the Terraform state with the retrieved space details
	d.Set("name", space.Name)
	d.Set("description", space.Desc)

	return diag.Diagnostics{}
}

func resourceSpaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	// Check if the description has changed
	if d.HasChange("description") {
		description := d.Get("description").(string)

		// Update the space description using the API
		_, err := client.UpdateSpace(spaceID, description)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error updating space: %w", err))
		}
	}

	// Re-read the resource to update the state
	return resourceSpaceRead(ctx, d, m)
}

func resourceSpaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	// Call the API to delete the space
	err := client.DeleteSpace(spaceID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting space: %w", err))
	}

	// Clear the resource ID to remove it from the state
	d.SetId("")

	return diag.Diagnostics{}
}
