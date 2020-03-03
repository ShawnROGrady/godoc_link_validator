package client

import (
	"net/http"
	"path"
)

// Client represents an http client.
type Client struct {
	baseURL string
	cl      *http.Client
}

// New constructs a new client
func New(baseURL string) Client {
	return Client{
		cl:      &http.Client{},
		baseURL: baseURL,
	}
}

// Get makes an http get request.
// Paths are joined, so if the client has
// baseURL https://example.com, c.Get("resource", "value")
// the corresponding GET call would be to https://example.com/resource/value
func (c *Client) Get(paths ...string) (*http.Response, error) {
	p := append([]string{c.baseURL}, paths...)
	url := path.Join(p...)
	return c.cl.Get(url)
}
