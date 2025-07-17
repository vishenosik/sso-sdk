package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/vishenosik/sso-sdk/api"
	"github.com/vishenosik/sso-sdk/gen/client"
)

func (c *Client) Ping() error {
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

func Ping2() error {

	cli, err := client.NewClient("http://localhost:8080")
	if err != nil {
		return err
	}

	resp, err := cli.GetApiSystemPing(context.TODO(), &client.GetApiSystemPingParams{
		Q: "hello",
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	log.Println("Response Body:", string(body))

	resp, err = cli.PostApiSystemMetricsLog(context.TODO(),
		[]client.ApiMetric{
			{
				ParamOne:   "param_one",
				ParamTwo:   "param_two",
				ParamThree: "param_three",
			},
		})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	log.Println("Response Body:", string(body))
	return nil
}
