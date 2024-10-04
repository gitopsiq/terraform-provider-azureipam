package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSpaces_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpacesConfig,
				Check: resource.ComposeTestCheckFunc(
					// Check that we have the expected number of spaces
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.#", "2"),
					
					// Check attributes of the first space
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.name", "space1"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.description", "This is space 1"),
					
					// Check attributes of the second space
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.name", "space2"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.description", "This is space 2"),
					
					// Use a custom check function to validate the spaces
					testAccCheckSpacesDataSource("data.azureipam_spaces.test"),
				),
			},
		},
	})
}

// Custom check function to validate the spaces data source
func testAccCheckSpacesDataSource(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		// Check for the presence of computed values
		attributes := rs.Primary.Attributes
		if v, ok := attributes["spaces.#"]; !ok || v != "2" {
			return fmt.Errorf("Expected 2 spaces, got %s", v)
		}

		spacesAttributes := []string{"name", "description"}
		for i := 0; i < 2; i++ {
			for _, attr := range spacesAttributes {
				if v, ok := attributes[fmt.Sprintf("spaces.%d.%s", i, attr)]; !ok || v == "" {
					return fmt.Errorf("Space %d is missing attribute %s", i, attr)
				}
			}
		}

		return nil
	}
}

const testAccDataSourceSpacesConfig = `
data "azureipam_spaces" "test" {}

output "all_spaces" {
  value = data.azureipam_spaces.test.spaces
}
`