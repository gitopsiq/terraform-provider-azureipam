package client

import (
    "net/http"
)

// Client is the structure for the Azure IPAM API client
type Client struct {
    BaseURL    string
    Token      string
    HTTPClient *http.Client
}

// NewClient initializes and returns a new API client
func NewClient(baseURL, token string) *Client {
    return &Client{
        BaseURL:    baseURL,
        Token:      token,
        HTTPClient: &http.Client{},
    }
}