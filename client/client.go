package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Client represents the Azure IPAM API client
type Client struct {
	Token      string
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new Azure IPAM API client with the given token and base URL
func NewClient(token, baseURL string) *Client {
	return &Client{
		Token:   token,
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetHTTPClient allows setting a custom HTTP client (useful for testing)
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.HTTPClient = httpClient
}

// makeRequest handles making an HTTP request with proper error handling and logging
func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	// Construct the URL for the API endpoint
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	// Marshal the request body, if provided
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Content-Type", "application/json")

	// Log the request for debugging purposes
	log.Printf("[DEBUG] %s request to %s", method, url)
	if reqBody != nil {
		log.Printf("[DEBUG] Request body: %s", reqBody)
	}

	// Send the HTTP request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	// Log the response status for debugging purposes
	log.Printf("[DEBUG] Response status: %d", resp.StatusCode)

	// Check for non-2xx status codes and return an error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-2xx response: %d, body: %s", resp.StatusCode, string(body))
	}

	return resp, nil
}

// CreateSpace creates a new space in the Azure IPAM API
func (c *Client) CreateSpace(name, desc string) (*Space, error) {
	space := Space{
		Name: name,
		Desc: desc,
	}

	// Make the API request
	resp, err := c.makeRequest("POST", "/api/spaces", space)
	if err != nil {
		return nil, fmt.Errorf("error creating space: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response into the Space struct
	var createdSpace Space
	if err := json.NewDecoder(resp.Body).Decode(&createdSpace); err != nil {
		return nil, fmt.Errorf("error decoding create space response: %w", err)
	}

	return &createdSpace, nil
}

// ListSpaces retrieves all spaces from the Azure IPAM API
func (c *Client) ListSpaces() ([]Space, error) {
	// Make the API request
	resp, err := c.makeRequest("GET", "/api/spaces", nil)
	if err != nil {
		return nil, fmt.Errorf("error listing spaces: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response into a slice of Space structs
	var spaces []Space
	if err := json.NewDecoder(resp.Body).Decode(&spaces); err != nil {
		return nil, fmt.Errorf("error decoding list spaces response: %w", err)
	}

	return spaces, nil
}

// GetSpace retrieves a specific space by its ID from the Azure IPAM API
func (c *Client) GetSpace(id string) (*Space, error) {
	// Make the API request
	resp, err := c.makeRequest("GET", fmt.Sprintf("/api/spaces/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("error retrieving space: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response into the Space struct
	var space Space
	if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
		return nil, fmt.Errorf("error decoding get space response: %w", err)
	}

	return &space, nil
}

// DeleteSpace deletes a specific space by its ID from the Azure IPAM API
func (c *Client) DeleteSpace(id string) error {
	// Make the API request
	resp, err := c.makeRequest("DELETE", fmt.Sprintf("/api/spaces/%s", id), nil)
	if err != nil {
		return fmt.Errorf("error deleting space: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
