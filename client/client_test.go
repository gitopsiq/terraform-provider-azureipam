package client_test

import (
    "context"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/gitopsiq/terraform-provider-azureipam/client"
)

// TestGetAllSpaces tests retrieving all spaces
func TestGetAllSpaces(t *testing.T) {
    // Mock API response
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`[{"id": "1", "name": "Test Space", "description": "A test space"}]`))
    }

    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()

    c := client.NewClient(server.URL, "fake-token")
    
    // Pass context with background
    spaces, err := c.GetAllSpaces(context.Background())

    assert.NoError(t, err)
    assert.Equal(t, 1, len(spaces))
    assert.Equal(t, "Test Space", spaces[0].Name)
}

// TestCreateSpace tests creating a space
func TestCreateSpace(t *testing.T) {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"id": "1", "name": "New Space", "description": "A new test space"}`))
    }

    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()

    c := client.NewClient(server.URL, "fake-token")

    space, err := c.CreateSpace(context.Background(), "New Space", "A new test space")

    assert.NoError(t, err)
    assert.Equal(t, "New Space", space.Name)
    assert.Equal(t, "A new test space", space.Description)
}

// TestGetSpace tests retrieving a single space by ID
func TestGetSpace(t *testing.T) {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"id": "1", "name": "Test Space", "description": "A test space"}`))
    }

    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()

    c := client.NewClient(server.URL, "fake-token")

    space, err := c.GetSpace(context.Background(), "1")

    assert.NoError(t, err)
    assert.Equal(t, "Test Space", space.Name)
    assert.Equal(t, "A test space", space.Description)
}

// TestUpdateSpace tests updating a space
func TestUpdateSpace(t *testing.T) {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"id": "1", "name": "Updated Space", "description": "An updated test space"}`))
    }

    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()

    c := client.NewClient(server.URL, "fake-token")

    space, err := c.UpdateSpace(context.Background(), "1", "Updated Space", "An updated test space")

    assert.NoError(t, err)
    assert.Equal(t, "Updated Space", space.Name)
    assert.Equal(t, "An updated test space", space.Description)
}

// TestDeleteSpace tests deleting a space
func TestDeleteSpace(t *testing.T) {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNoContent)
    }

    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()

    c := client.NewClient(server.URL, "fake-token")

    err := c.DeleteSpace(context.Background(), "1")

    assert.NoError(t, err)
}

// Table-Driven Test for GetAllSpaces with different scenarios
func TestGetAllSpaces_Scenarios(t *testing.T) {
    tests := []struct {
        name           string
        responseStatus int
        responseBody   string
        expectedError  bool
        expectedLength int
    }{
        {
            name:           "Success",
            responseStatus: http.StatusOK,
            responseBody:   `[{"id": "1", "name": "Test Space", "description": "A test space"}]`,
            expectedError:  false,
            expectedLength: 1,
        },
        {
            name:           "Empty Response",
            responseStatus: http.StatusOK,
            responseBody:   `[]`,
            expectedError:  false,
            expectedLength: 0,
        },
        {
            name:           "Invalid Response",
            responseStatus: http.StatusOK,
            responseBody:   `invalid`,
            expectedError:  true,
            expectedLength: 0,
        },
        {
            name:           "Server Error",
            responseStatus: http.StatusInternalServerError,
            responseBody:   ``,
            expectedError:  true,
            expectedLength: 0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            handler := func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(tt.responseStatus)
                w.Write([]byte(tt.responseBody))
            }

            server := httptest.NewServer(http.HandlerFunc(handler))
            defer server.Close()

            c := client.NewClient(server.URL, "fake-token")
            spaces, err := c.GetAllSpaces(context.Background())

            if tt.expectedError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expectedLength, len(spaces))
            }
        })
    }
}
