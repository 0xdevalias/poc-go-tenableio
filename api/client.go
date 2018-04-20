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

	Workbenches *WorkbenchesAPI
}

type WorkbenchesAPI struct {
	client *Client

	commonFilterHelpers
}

func DefaultClient(accessKey string, secretKey string) *Client {
	c := resty.DefaultClient
	c.SetHostURL(BaseUrl)
	c.SetHeader("X-ApiKeys", fmt.Sprintf("accessKey=%s; secretKey=%s", accessKey, secretKey))

	client := Client{restyClient: c}
	client.Workbenches = &WorkbenchesAPI{client: &client}

	return &client
}

func (c *Client) WithDebug() *Client {
	c.restyClient.SetDebug(true)
	return c
}
