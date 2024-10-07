package client

import (
    "fmt"
    "bytes"
    "encoding/json"
    "net/http"
    "context"
)

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
    url := c.BaseURL + path

    var reqBody *bytes.Buffer
    if body != nil {
        jsonBody, err := json.Marshal(body)
        if err != nil {
            return nil, err
        }
        reqBody = bytes.NewBuffer(jsonBody)
    } else {
        reqBody = &bytes.Buffer{}
    }

    req, err := http.NewRequest(method, url, reqBody)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", "Bearer "+c.Token)
    req.Header.Set("Content-Type", "application/json")

    return req, nil
}

// checkStatusCode is a helper function to verify that the response status code is as expected
func checkStatusCode(resp *http.Response, expectedStatusCode int) error {
    if resp.StatusCode != expectedStatusCode {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }
    return nil
}

// newRequestWithContext creates an HTTP request with context, adding headers as needed
func (c *Client) newRequestWithContext(ctx context.Context, method, urlStr string, body *bytes.Buffer) (*http.Request, error) {
    req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+urlStr, body)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", "Bearer "+c.Token) // Assuming c.Token holds the authentication token
    return req, nil
}
