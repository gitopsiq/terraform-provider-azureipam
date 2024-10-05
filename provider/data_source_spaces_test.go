package provider

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSpaces_basic(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && r.URL.Path == "/api/spaces" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				{"id": "space-123", "name": "test-space", "description": "Test description"},
				{"id": "space-456", "name": "another-space", "description": "Another description"}
			]`))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer mockServer.Close()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpacesConfig(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.#", "2"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.0.name", "test-space"),
					resource.TestCheckResourceAttr("data.azureipam_spaces.test", "spaces.1.name", "another-space"),
				),
			},
		},
	})
}

func testAccDataSourceSpacesConfig(apiURL string) string {
	return `
provider "azureipam" {
  api_endpoint = "` + apiURL + `"
  token        = "fake-token"
}

data "azureipam_spaces" "test" {}
`
}
