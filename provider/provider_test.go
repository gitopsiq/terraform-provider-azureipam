package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories is used to instantiate a provider during acceptance testing.
var providerFactories = map[string]func() (*schema.Provider, error){
	"azureipam": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
