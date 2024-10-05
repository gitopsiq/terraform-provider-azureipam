package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Space struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

type Client struct {
	Token      string
	ApiEndpoint string
}

func NewClient(token, apiEndpoint string) *Client {
	return &Client{
		Token:      token,
		ApiEndpoint: apiEndpoint,
	}
}

// CreateSpace creates a new space
func (c *Client) CreateSpace(name, description string) (*Space, error) {
	url := fmt.Sprintf("%s/api/spaces", c.ApiEndpoint)

	spaceData := map[string]string{
		"name":        name,
		"description": description,
	}

	// Serialize request data
	body, err := json.Marshal(spaceData)
	if err != nil {
		return nil, err
	}

	// Make HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("received non-2xx response: %d", resp.StatusCode)
	}

	// Parse the response
	var createdSpace Space
	if err := json.NewDecoder(resp.Body).Decode(&createdSpace); err != nil {
		return nil, err
	}

	// Log the created space response
	log.Printf("[DEBUG] API returned created space: ID=%s, name=%s, description=%s", createdSpace.ID, createdSpace.Name, createdSpace.Desc)

	return &createdSpace, nil
}

// GetSpace fetches a space by ID
func (c *Client) GetSpace(spaceID string) (*Space, error) {
	url := fmt.Sprintf("%s/api/spaces/%s", c.ApiEndpoint, spaceID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-2xx response: %d", resp.StatusCode)
	}

	// Parse the response
	var space Space
	if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
		return nil, err
	}

	// Log the retrieved space response
	log.Printf("[DEBUG] API returned space: ID=%s, name=%s, description=%s", space.ID, space.Name, space.Desc)

	return &space, nil
}
