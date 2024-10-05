package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSpaces_basic(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock the GET /api/spaces response
		if r.Method == "GET" && r.URL.Path == "/api/spaces" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				{
					"id": "space-123",
					"name": "test-space",
					"desc": "Test description"
				},
				{
					"id": "space-456",
					"name": "another-space",
					"desc": "Another test description"
				}
			]`))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer mockServer.Close()

	// Define the Terraform test case
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders(mockServer.URL),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpacesConfig(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.#", "2"), // Ensure we get 2 spaces
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.id", "space-123"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.name", "test-space"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.desc", "Test description"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.id", "space-456"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.name", "another-space"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.desc", "Another test description"),
				),
			},
		},
	})
}

func testAccDataSourceSpacesConfig(apiURL string) string {
	return fmt.Sprintf(`
provider "azureipam" {
  api_endpoint = "%s"
  token        = "fake-token"
}

data "azureipam_spaces" "test" {}
`, apiURL)
}
