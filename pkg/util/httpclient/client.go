package httpclient

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
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
	debug(httputil.DumpRequestOut(req, true))
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	debug(httputil.DumpResponse(res, true))
	return res, err
}

func (c Client) Post(ctx context.Context, url string, header http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	debug(httputil.DumpRequestOut(req, true))
	if err != nil {
		return nil, err
	}

	if header != nil {
		req.Header = header
	}

	res, err := c.client.Do(req)
	debug(httputil.DumpResponse(res, true))
	return res, err
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

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("\n%s", data)
	} else {
		log.Fatalf("\n%s", err)
	}
}
