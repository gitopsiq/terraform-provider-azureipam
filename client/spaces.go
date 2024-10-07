package client

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"
)

// CreateSpace creates a new IPAM space
func (c *Client) CreateSpace(ctx context.Context, name string, desc string) (*Space, error) {
    spaceReq := Space{
        Name:        name,
        Description: desc,
    }

    reqBody, err := json.Marshal(spaceReq)
    if err != nil {
        return nil, err
    }

    req, err := c.newRequestWithContext(ctx, "POST", "/api/spaces", bytes.NewBuffer(reqBody))
    if err != nil {
        return nil, err
    }

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Check for the expected status code using the utility function
    if err := checkStatusCode(resp, http.StatusCreated); err != nil {
        return nil, err
    }

    var space Space
    if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
        return nil, err
    }

    return &space, nil
}

// GetSpace retrieves a space by its ID
func (c *Client) GetSpace(ctx context.Context, id string) (*Space, error) {
    req, err := c.newRequestWithContext(ctx, "GET", fmt.Sprintf("/api/spaces/%s", id), nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Check if the space was not found
    if resp.StatusCode == http.StatusNotFound {
        return nil, fmt.Errorf("space not found")
    }

    if err := checkStatusCode(resp, http.StatusOK); err != nil {
        return nil, err
    }

    var space Space
    if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
        return nil, err
    }

    return &space, nil
}

// UpdateSpace updates an existing IPAM space
func (c *Client) UpdateSpace(ctx context.Context, id string, name string, desc string) (*Space, error) {
    spaceReq := Space{
        Name:        name,
        Description: desc,
    }

    reqBody, err := json.Marshal(spaceReq)
    if err != nil {
        return nil, err
    }

    req, err := c.newRequestWithContext(ctx, "PATCH", fmt.Sprintf("/api/spaces/%s", id), bytes.NewBuffer(reqBody))
    if err != nil {
        return nil, err
    }

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if err := checkStatusCode(resp, http.StatusOK); err != nil {
        return nil, err
    }

    var space Space
    if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
        return nil, err
    }

    return &space, nil
}

// DeleteSpace deletes a space by its ID
func (c *Client) DeleteSpace(ctx context.Context, id string) error {
    req, err := c.newRequestWithContext(ctx, "DELETE", fmt.Sprintf("/api/spaces/%s", id), nil)
    if err != nil {
        return err
    }

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if err := checkStatusCode(resp, http.StatusNoContent); err != nil {
        return err
    }

    return nil
}

// GetAllSpaces retrieves all spaces (with pagination handling, if needed)
func (c *Client) GetAllSpaces(ctx context.Context) ([]Space, error) {
    var allSpaces []Space
    page := 1

    for {
        req, err := c.newRequestWithContext(ctx, "GET", fmt.Sprintf("/api/spaces?page=%d", page), nil)
        if err != nil {
            return nil, err
        }

        resp, err := c.HTTPClient.Do(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        if err := checkStatusCode(resp, http.StatusOK); err != nil {
            return nil, err
        }

        var spaces []Space
        if err := json.NewDecoder(resp.Body).Decode(&spaces); err != nil {
            return nil, err
        }

        if len(spaces) == 0 {
            break // No more spaces
        }

        allSpaces = append(allSpaces, spaces...)
        page++
    }

    return allSpaces, nil
}
