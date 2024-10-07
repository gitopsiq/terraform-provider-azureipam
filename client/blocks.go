package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetBlocks retrieves all blocks within a space
func (c *Client) GetBlocks(spaceID string) ([]Block, error) {
    path := fmt.Sprintf("/api/spaces/%s/blocks", spaceID)
    req, err := c.newRequest("GET", path, nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to get blocks: status code %d", resp.StatusCode)
    }

    var blocks []Block
    err = json.NewDecoder(resp.Body).Decode(&blocks)
    if err != nil {
        return nil, err
    }

    return blocks, nil
}