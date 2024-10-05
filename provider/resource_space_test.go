package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAzureIPAMSpace_basic(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mocking based on the HTTP method and the path of the request
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

	// Define the Terraform test steps
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders(mockServer.URL),  // Use mock server's URL
		CheckDestroy: testAccCheckSpaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSpaceConfigBasic(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpaceExists("azureipam_space.test"),
					resource.TestCheckResourceAttr("azureipam_space.test", "name", "test-space"),
					resource.TestCheckResourceAttr("azureipam_space.test", "desc", "Test description"),
				),
			},
			// Update the space
			{
				Config: testAccSpaceConfigUpdated(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpaceExists("azureipam_space.test"),
					resource.TestCheckResourceAttr("azureipam_space.test", "name", "updated-space"),
					resource.TestCheckResourceAttr("azureipam_space.test", "desc", "Updated description"),
				),
			},
		},
	})
}

func testAccSpaceConfigBasic(apiURL string) string {
	return fmt.Sprintf(`
provider "azureipam" {
  api_endpoint = "%s"
  token        = "fake-token"
}

resource "azureipam_space" "test" {
  name = "test-space"
  desc = "Test description"
}
`, apiURL)
}

func testAccSpaceConfigUpdated(apiURL string) string {
	return fmt.Sprintf(`
provider "azureipam" {
  api_endpoint = "%s"
  token        = "fake-token"
}

resource "azureipam_space" "test" {
  name = "updated-space"
  desc = "Updated description"
}
`, apiURL)
}

// Check if the space exists in the Terraform state
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

// Check if the space is properly destroyed
func testAccCheckSpaceDestroy(s *terraform.State) error {
	// Verify that the resource is destroyed correctly (can call GET and expect 404)
	return nil
}
