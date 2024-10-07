package client

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// CreateReservation creates a CIDR reservation within a block
func (c *Client) CreateReservation(spaceID, blockID string, cidr string) (*Reservation, error) {
    path := fmt.Sprintf("/api/spaces/%s/blocks/%s/reservations", spaceID, blockID)
    body := map[string]interface{}{
        "cidr": cidr,
    }

    req, err := c.newRequest("POST", path, body)
    if err != nil {
        return nil, err
    }

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return nil, fmt.Errorf("failed to create reservation: status code %d", resp.StatusCode)
    }

    var reservation Reservation
    err = json.NewDecoder(resp.Body).Decode(&reservation)
    if err != nil {
        return nil, err
    }

    return &reservation, nil
}