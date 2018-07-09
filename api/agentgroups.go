package api

import (
	"strconv"
)

type AgentGroupRef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AgentGroupName struct {
	Name string `json:"name"`
}

// AgentGroups
// Ref: https://cloud.tenable.com/api#/resources/agent-groups/list
type AgentGroups struct {
	Groups []AgentGroup `json:"groups"`
}

// AgentGroup contains the details of an agent group.
// Not all fields may be returned unless Agent Enhancements are enabled for your account.
// Ref: https://cloud.tenable.com/api#/resources/agent-groups
type AgentGroup struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`

	CreationDate         int `json:"creation_date"`
	LastModificationDate int `json:"last_modification_date"`

	Name string `json:"name"`

	OwnerID   int    `json:"owner_id"`
	OwnerUUID string `json:"owner_uuid"`
	Owner     string `json:"owner"`
	OwnerName string `json:"owner_name"`

	Agents []Agent `json:"agents"`

	Pagination Pagination `json:"pagination"`

	Shared          int `json:"shared"`
	UserPermissions int `json:"user_permissions"`
}

// List returns the agent groups for the given scanner.
//   Ref: https://cloud.tenable.com/api#/resources/agent-groups/list
func (c *AgentGroupsAPI) List(scannerID int) ([]AgentGroup, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
		}).SetResult(AgentGroups{})

	r, err := req.Get("/scanners/{scanner_id}/agent-groups")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*AgentGroups).Groups, nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Create creates an agent group on the given scanner.
//   Ref: https://cloud.tenable.com/api#/resources/agent-groups/create
func (c *AgentGroupsAPI) Create(scannerID int, groupName string) (*AgentGroup, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
		}).SetBody(AgentGroupName{Name: groupName}).
		SetResult(AgentGroup{})

	r, err := req.Post("/scanners/{scanner_id}/agent-groups")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*AgentGroup), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Configure changes the name of the given agent group.
//   Ref: https://cloud.tenable.com/api#/resources/agent-groups/configure
func (c *AgentGroupsAPI) Configure(scannerID int, groupID int, groupName string) (bool, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
			"group_id":   strconv.FormatInt(int64(groupID), 10),
		}).SetBody(AgentGroupName{Name: groupName})

	r, err := req.Put("/scanners/{scanner_id}/agent-groups/{group_id}")
	if err != nil {
		return false, err
	}

	switch r.StatusCode() {
	case 200:
		return true, nil
	default:
		return false, errorResponseFormatter(r)
	}
}

// Details returns details for the given agent group.
// Agent records which belong to this group will also be returned.
// You can apply filtering, sorting, or pagination to the agent records.
//   Ref: https://cloud.tenable.com/api#/resources/agent-groups/details
func (c *AgentGroupsAPI) Details(scannerID int, groupID int) (*AgentGroup, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
			"group_id":   strconv.FormatInt(int64(groupID), 10),
		}).SetResult(AgentGroup{})

	// TODO: Implement filtering/etc optional query params

	r, err := req.Get("/scanners/{scanner_id}/agent-groups/{group_id}")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*AgentGroup), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Delete deletes an agent group from the given scanner.
//   Ref: https://cloud.tenable.com/api#/resources/agent-groups/delete
func (c *AgentGroupsAPI) Delete(scannerID int, groupID int) (bool, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
			"group_id":   strconv.FormatInt(int64(groupID), 10),
		})

	r, err := req.Delete("/scanners/{scanner_id}/agent-groups/{group_id}")
	if err != nil {
		return false, err
	}

	switch r.StatusCode() {
	case 200:
		return true, nil
	default:
		return false, errorResponseFormatter(r)
	}
}

// AddAgent adds an agent to the given agent group.
//   Ref: https://cloud.tenable.com/api#/resources/agent-groups/add-agent
func (c *AgentGroupsAPI) AddAgent(scannerID int, groupID int, agentID int) (bool, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
			"group_id":   strconv.FormatInt(int64(groupID), 10),
			"agent_id":   strconv.FormatInt(int64(agentID), 10),
		})

	r, err := req.Put("/scanners/{scanner_id}/agent-groups/{group_id}/agents/{agent_id}")
	if err != nil {
		return false, err
	}

	switch r.StatusCode() {
	case 200:
		return true, nil
	default:
		return false, errorResponseFormatter(r)
	}
}

// DeleteAgent deletes an agent from the given agent group.
//   Ref: https://cloud.tenable.com/api#/resources/agent-groups/delete-agent
func (c *AgentGroupsAPI) DeleteAgent(scannerID int, groupID int, agentID int) (bool, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
			"group_id":   strconv.FormatInt(int64(groupID), 10),
			"agent_id":   strconv.FormatInt(int64(agentID), 10),
		})

	r, err := req.Delete("/scanners/{scanner_id}/agent-groups/{group_id}/agents/{agent_id}")
	if err != nil {
		return false, err
	}

	switch r.StatusCode() {
	case 200:
		return true, nil
	default:
		return false, errorResponseFormatter(r)
	}
}
