package provider

import (
	"context"
	"fmt"
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
				Description: "A description of the space",
			},
		},
	}
}

func resourceSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceName := d.Get("name").(string)
	spaceDesc := d.Get("description").(string)

	// Make API request to create a space
	createdSpace, err := client.CreateSpace(spaceName, spaceDesc)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating space: %s", err))
	}

	// Set the resource ID
	d.SetId(createdSpace.ID)

	// Save name and description to state
	d.Set("name", createdSpace.Name)
	d.Set("description", createdSpace.Desc)

	return resourceSpaceRead(ctx, d, m)
}

func resourceSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	// Make API request to read space details
	space, err := client.GetSpace(spaceID)
	if err != nil {
		if space == nil {
			// Resource no longer exists
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	// Save the space data to state
	d.Set("name", space.Name)
	d.Set("description", space.Desc)

	return diag.Diagnostics{}
}

func resourceSpaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*client.Client)

	spaceID := d.Id()

	if d.HasChange("description") {
		description := d.Get("description").(string)

		// Make API request to update space description
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

	// Make API request to delete the space
	err := client.DeleteSpace(spaceID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting space: %s", err))
	}

	// Remove the resource from state
	d.SetId("")

	return diag.Diagnostics{}
}
