package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/vishenosik/sso-sdk/api"
)

func (c *client) Ping() error {
	request, err := c.newRequest(http.MethodGet, api.PingMethod, nil)
	if err != nil {
		return err
	}

	res, err := c.cli.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
