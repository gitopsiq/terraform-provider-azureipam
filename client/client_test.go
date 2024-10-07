package client_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/gitopsiq/terraform-provider-azureipam/client"
)

func TestGetSpaces(t *testing.T) {
    // Mock API response
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`[{"id": "1", "name": "Test Space", "description": "A test space"}]`))
    }

    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()

    c := client.NewClient(server.URL, "fake-token")
    spaces, err := c.GetSpaces()

    assert.NoError(t, err)
    assert.Equal(t, 1, len(spaces))
    assert.Equal(t, "Test Space", spaces[0].Name)
}