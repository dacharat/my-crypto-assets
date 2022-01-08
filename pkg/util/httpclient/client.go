package httpclient

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

//go:generate mockgen -source=./client.go -destination=./mock_client/mock_client.go -package=mock_client
type IClient interface {
	Get(ctx context.Context, url string, header http.Header, opts ...Option) (*http.Response, error)
	Post(ctx context.Context, url string, header http.Header, body io.Reader, opts ...Option) (*http.Response, error)
	Put(ctx context.Context, url string, header http.Header, body io.Reader, opts ...Option) (*http.Response, error)
}

func NewClient() IClient {
	return &Client{
		client: http.DefaultClient,
	}
}

type Client struct {
	client *http.Client
}

func (c Client) Get(ctx context.Context, url string, header http.Header, opts ...Option) (*http.Response, error) {
	return c.do(ctx, http.MethodGet, url, header, nil, opts...)
}

func (c Client) Post(ctx context.Context, url string, header http.Header, body io.Reader, opts ...Option) (*http.Response, error) {
	return c.do(ctx, http.MethodPost, url, header, body, opts...)
}

func (c Client) Put(ctx context.Context, url string, header http.Header, body io.Reader, opts ...Option) (*http.Response, error) {
	return c.do(ctx, http.MethodPut, url, header, body, opts...)
}

func (c *Client) do(ctx context.Context, method string, url string, header http.Header, body io.Reader, opts ...Option) (*http.Response, error) {
	ho := &httpOption{}
	for _, opt := range opts {
		opt(ho)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	// log before set header
	if !ho.withoutReqLog {
		debug(httputil.DumpRequestOut(req, true))
	}

	if header != nil {
		req.Header = header
	}

	res, err := c.client.Do(req)

	if !ho.withoutResLog {
		debug(httputil.DumpResponse(res, true))
	}

	return res, err
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("\n%s", data)
	} else {
		log.Fatalf("\n%s", err)
	}
}
