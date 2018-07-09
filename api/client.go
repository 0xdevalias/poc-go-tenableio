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

	Agents      *AgentsAPI
	AgentGroups *AgentGroupsAPI
	Plugins     *PluginsAPI
	Workbenches *WorkbenchesAPI
}

type AgentsAPI struct{ client *Client }

type AgentGroupsAPI struct{ client *Client }

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
	client.Agents = &AgentsAPI{client: &client}
	client.AgentGroups = &AgentGroupsAPI{client: &client}
	client.Plugins = &PluginsAPI{client: &client}
	client.Workbenches = &WorkbenchesAPI{client: &client}

	return &client
}

func (c *Client) WithDebug() *Client {
	c.restyClient.SetDebug(true)
	return c
}
