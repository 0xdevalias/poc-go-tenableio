package api

import "fmt"

func (c *WorkbenchesAPI) WithFilters(filters ...Filter) *WorkbenchesAPI {
	return &WorkbenchesAPI{
		client:              c.client,
		commonFilterHelpers: c.commonFilterHelpers.WithFilters(filters...),
	}
}

// A list of up to 5000 assets. The list can be modified using filters.
//   Ref: https://cloud.tenable.com/api#/resources/workbenches/assets
func (c *WorkbenchesAPI) Assets() (*Assets, error) {
	req := c.client.restyClient.R().
		SetResult(Assets{})

	c.applyCommonFilters(req)

	r, err := req.Get("/workbenches/assets")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*Assets), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// A list of up to 5000 of the vulnerabilities recorded for a given asset. By default, this list is sorted by vulnerability count, descending.
//   Ref: https://cloud.tenable.com/api#/resources/workbenches/asset-vulnerabilities
func (c *WorkbenchesAPI) AssetVulns(assetID string) (*AssetVulnerabilities, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{"asset_id": assetID}).
		SetResult(AssetVulnerabilities{})

	c.applyCommonFilters(req)

	r, err := req.Get("/workbenches/assets/{asset_id}/vulnerabilities")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*AssetVulnerabilities), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Get the details for a vulnerability recorded on a given asset.
//   Ref: https://cloud.tenable.com/api#/resources/workbenches/asset-vulnerability-info
func (c *WorkbenchesAPI) AssetVulnInfo(assetID string, pluginID int) (*AssetVulnInfo, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"asset_id":  assetID,
			"plugin_id": fmt.Sprintf("%v", pluginID),
		}).
		SetResult(AssetVulnInfo{})

	c.applyCommonFilters(req)

	r, err := req.Get("/workbenches/assets/{asset_id}/vulnerabilities/{plugin_id}/info")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*AssetVulnInfo), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Get the vulnerability outputs for a plugin recorded on a given asset.
//   Ref: https://cloud.tenable.com/api#/resources/workbenches/asset-vulnerability-output
func (c *WorkbenchesAPI) AssetVulnOutput(assetID string, pluginID int) (*AssetVulnOutputs, error) {
	req := c.client.restyClient.R().
		SetPathParams(map[string]string{
			"asset_id":  assetID,
			"plugin_id": fmt.Sprintf("%v", pluginID),
		}).
		SetResult(AssetVulnOutputs{})

	c.applyCommonFilters(req)

	r, err := req.Get("/workbenches/assets/{asset_id}/vulnerabilities/{plugin_id}/outputs")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*AssetVulnOutputs), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}
