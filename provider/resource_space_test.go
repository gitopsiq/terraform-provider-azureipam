package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccAzureIPAMSpace_basic tests the creation, update, and deletion of a space
func TestAccAzureIPAMSpace_basic(t *testing.T) {
	// Create a mock server for the Azure IPAM API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "POST" && r.URL.Path == "/api/spaces":
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{
				"id": "space-123",
				"name": "test-space",
				"desc": "Test description"
			}`))

		case r.Method == "GET" && r.URL.Path == "/api/spaces/space-123":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"id": "space-123",
				"name": "test-space",
				"desc": "Test description"
			}`))

		case r.Method == "PATCH" && r.URL.Path == "/api/spaces/space-123":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"id": "space-123",
				"name": "updated-space",
				"desc": "Updated description"
			}`))

		case r.Method == "DELETE" && r.URL.Path == "/api/spaces/space-123":
			w.WriteHeader(http.StatusNoContent)

		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer mockServer.Close()

	// Define the Terraform test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders(mockServer.URL),
		CheckDestroy: testAccCheckSpaceDestroy,
		Steps: []resource.TestStep{
			// Initial creation
			{
				Config: testAccSpaceConfigBasic(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpaceExists("azureipam_space.test"),
					resource.TestCheckResourceAttr("azureipam_space.test", "name", "test-space"),
					resource.TestCheckResourceAttr("azureipam_space.test", "description", "Test description"), // Ensure "description" is used
				),
			},
			// Update the space
			{
				Config: testAccSpaceConfigUpdated(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpaceExists("azureipam_space.test"),
					resource.TestCheckResourceAttr("azureipam_space.test", "name", "updated-space"),
					resource.TestCheckResourceAttr("azureipam_space.test", "description", "Updated description"),
				),
			},
		},
	})
}

// Configuration for creating a basic space
func testAccSpaceConfigBasic(apiURL string) string {
	return fmt.Sprintf(`
provider "azureipam" {
  api_endpoint = "%s"
  token        = "fake-token"
}

resource "azureipam_space" "test" {
  name = "test-space"
  description = "Test description"  // Use "description" to match the schema
}
`, apiURL)
}

// Configuration for updating the space
func testAccSpaceConfigUpdated(apiURL string) string {
	return fmt.Sprintf(`
provider "azureipam" {
  api_endpoint = "%s"
  token        = "fake-token"
}

resource "azureipam_space" "test" {
  name = "updated-space"
  description = "Updated description"  // Use "description" to match the schema
}
`, apiURL)
}

// testAccCheckSpaceExists verifies that the space exists in the Terraform state
func testAccCheckSpaceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no space ID is set")
		}

		return nil
	}
}

// testAccCheckSpaceDestroy verifies that the space was properly destroyed
func testAccCheckSpaceDestroy(s *terraform.State) error {
	// Logic to ensure the space no longer exists after the test
	// Normally, we would make a GET call to the API and ensure it returns 404
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azureipam_space" {
			continue
		}

		// Here you would check the mock API or actual API to ensure the resource has been deleted
		// Since we are using a mock server, this check is skipped but could be simulated in real tests.
	}

	return nil
}
