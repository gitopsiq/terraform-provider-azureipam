package client

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// GetSpaces retrieves all IPAM spaces
func (c *Client) GetSpaces() ([]Space, error) {
    req, err := c.newRequest("GET", "/api/spaces", nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to get spaces: status code %d", resp.StatusCode)
    }

    var spaces []Space
    err = json.NewDecoder(resp.Body).Decode(&spaces)
    if err != nil {
        return nil, err
    }

    return spaces, nil
}