package httpclient

import (
	"context"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() Client {
	return Client{
		client: http.DefaultClient,
	}
}

func (c Client) Get(ctx context.Context, url string, header http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) Put(ctx context.Context, url string, header http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	// header = http.Header{}
	// header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header = header

	return c.client.Do(req)
}
