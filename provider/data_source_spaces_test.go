package provider

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSpaces_basic(t *testing.T) {
	// Create a mock server that simulates the Azure IPAM API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock the GET /api/spaces response
		if r.Method == "GET" && r.URL.Path == "/api/spaces" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				{
					"name": "test-space",
					"description": "Test description"
				},
				{
					"name": "another-space",
					"description": "Another test description"
				}
			]`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Unauthorized"}`))
		}
	}))
	defer mockServer.Close()

	// Define the Terraform test case
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders(mockServer.URL),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpacesConfig(),  // No provider block here
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.#", "2"),  // Ensure we get 2 spaces
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.name", "test-space"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.description", "Test description"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.name", "another-space"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.description", "Another test description"),
				),
			},
		},
	})
}

func testAccDataSourceSpacesConfig() string {
	return `
data "azureipam_spaces" "test" {}
`
}
