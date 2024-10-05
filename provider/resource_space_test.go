package provider

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureIPAMSpace_basic(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "POST" && r.URL.Path == "/api/spaces":
			// Ensure the mock API returns the correct description when creating the space
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"id": "space-123", "name": "test-space", "description": "Test description"}`))
		case r.Method == "GET" && r.URL.Path == "/api/spaces/space-123":
			// Ensure the mock API returns the correct description when reading the space
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id": "space-123", "name": "test-space", "description": "Test description"}`))
		case r.Method == "DELETE" && r.URL.Path == "/api/spaces/space-123":
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer mockServer.Close()

	// Define the Terraform test case
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSpaceConfigBasic(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("azureipam_space.test", "name", "test-space"),
					resource.TestCheckResourceAttr("azureipam_space.test", "description", "Test description"),  // Check the description
				),
			},
		},
	})
}

func testAccSpaceConfigBasic(apiURL string) string {
	return `
provider "azureipam" {
  api_endpoint = "` + apiURL + `"
  token        = "fake-token"
}

resource "azureipam_space" "test" {
  name        = "test-space"
  description = "Test description"
}
`
}
