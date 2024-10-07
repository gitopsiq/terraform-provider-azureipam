package acceptance

import (
    "fmt"
    "os"
    "testing"

    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/gitopsiq/terraform-provider-azureipam/provider"
)

var testAccProviders = map[string]*schema.Provider{
    "azureipam": provider.Provider(), // Linking the provider to be used in the test
}

func TestAccAzureIPAMSpace_basic(t *testing.T) {
    resourceName := "azureipam_space.test_space"

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,    // Providers used during the test
        CheckDestroy: testAccCheckAzureIPAMSpaceDestroy, // Checks if the space is destroyed properly
        Steps: []resource.TestStep{
            {
                Config: testAccAzureIPAMSpaceConfig_basic(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    testAccCheckAzureIPAMSpaceExists(resourceName),
                    resource.TestCheckResourceAttr(resourceName, "name", "Test Space"),
                    resource.TestCheckResourceAttr(resourceName, "description", "A test space for acceptance testing"),
                ),
            },
        },
    })
}

func testAccPreCheck(t *testing.T) {
    if os.Getenv("AZURE_IPAM_TOKEN") == "" {
        t.Fatal("AZURE_IPAM_TOKEN must be set for acceptance tests")
    }
}

func testAccCheckAzureIPAMSpaceDestroy(s *terraform.State) error {
    // Implement any cleanup or validation that the resource was properly destroyed in the backend
    return nil
}

func testAccCheckAzureIPAMSpaceExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Not found: %s", resourceName)
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("No space ID is set")
        }

        return nil
    }
}

func testAccAzureIPAMSpaceConfig_basic() string {
    return `
resource "azureipam_space" "test_space" {
  name        = "Test Space"
  description = "A test space for acceptance testing"
}
`
}
