package api

import "strconv"

type AgentsList struct {
	Agents     []Agent    `json:"agents"`
	Pagination Pagination `json:"pagination"`
}

// List returns the agent list for the given scanner.
//   Ref: https://cloud.tenable.com/api#/resources/agents/list
func (c *AgentsAPI) List(scannerID int) (*AgentsList, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
		}).SetResult(AgentsList{})

	r, err := req.Get("/scanners/{scanner_id}/agents")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*AgentsList), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Get returns the specified agent for the given scanner.
//   Ref: https://cloud.tenable.com/api#/resources/agents/get
func (c *AgentsAPI) Get(scannerID int, agentID int) (*Agent, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
			"agent_id":   strconv.FormatInt(int64(agentID), 10),
		}).SetResult(Agent{})

	r, err := req.Get("/scanners/{scanner_id}/agents/{agent_id}")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*Agent), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Delete deletes an agent.
//   Ref: https://cloud.tenable.com/api#/resources/agents/delete
func (c *AgentsAPI) Delete(scannerID int, agentID int) (bool, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"scanner_id": strconv.FormatInt(int64(scannerID), 10),
			"agent_id":   strconv.FormatInt(int64(agentID), 10),
		}).SetResult(Agent{})

	r, err := req.Delete("/scanners/{scanner_id}/agents/{agent_id}")
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
