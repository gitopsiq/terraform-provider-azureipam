package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceSpace_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckSpaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSpaceConfig("test_space"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpaceExists("azureipam_space.test"),
					resource.TestCheckResourceAttr("azureipam_space.test", "name", "test_space"),
					resource.TestCheckResourceAttr("azureipam_space.test", "description", "Test space"),
				),
			},
			// Add more steps here to test updates, etc.
		},
	})
}

func testAccResourceSpaceConfig(name string) string {
	return fmt.Sprintf(`
resource "azureipam_space" "test" {
  name        = "%s"
  description = "Test space"
}
`, name)
}

func testAccCheckSpaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Space ID is set")
		}

		// Add code here to check if the space exists in your API

		return nil
	}
}

func testAccCheckSpaceDestroy(s *terraform.State) error {
	// Add code here to check that the space has been destroyed in your API
	return nil
}