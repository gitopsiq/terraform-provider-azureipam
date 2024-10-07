package client

import (
    "bytes"
    "encoding/json"
    "net/http"
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