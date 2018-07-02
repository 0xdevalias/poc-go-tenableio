package api

import (
	"fmt"

	"github.com/go-resty/resty"
)

const (
	BaseUrl string = "https://cloud.tenable.com/"
)

type Client struct {
	restyClient *resty.Client

	Plugins     *PluginsAPI
	Workbenches *WorkbenchesAPI
}

type PluginsAPI struct{ client *Client }

type WorkbenchesAPI struct {
	client *Client

	CommonFilterHelpers
}

func DefaultClient(accessKey string, secretKey string) *Client {
	c := resty.New()
	c.SetHostURL(BaseUrl)
	c.SetHeader("X-ApiKeys", fmt.Sprintf("accessKey=%s; secretKey=%s", accessKey, secretKey))

	client := Client{restyClient: c}
	client.Plugins = &PluginsAPI{client: &client}
	client.Workbenches = &WorkbenchesAPI{client: &client}

	return &client
}

func (c *Client) WithDebug() *Client {
	c.restyClient.SetDebug(true)
	return c
}
