package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/vishenosik/gocherry/pkg/config"
)

type Client struct {
	cli    *http.Client
	config Config
}

type Config struct {
	Server config.Server
}

func NewClient(config Config) *Client {
	return &Client{
		cli: &http.Client{
			Timeout: config.Server.Timeout,
		},
		config: config,
	}
}

func (c *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	const prefix = "api"
	return http.NewRequest(
		method,
		fmt.Sprintf("%s/%s/%s", c.config.Server, prefix, url),
		body,
	)
}
