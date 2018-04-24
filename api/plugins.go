package api

import "fmt"

// Returns the list of plugin families.
//   Ref: https://cloud.tenable.com/api#/resources/plugins/families
func (c *PluginsAPI) Families() (*PluginFamilies, error) {
	r, err := c.client.restyClient.R().
		SetResult(PluginFamilies{}).
		Get("/plugins/families")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*PluginFamilies), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Returns the list of plugins in a family.
//   Ref:https://cloud.tenable.com/api#/resources/plugins/family-details
func (c *PluginsAPI) FamilyDetails(familyID int) (*PluginFamilyDetails, error) {
	r, err := c.client.restyClient.R().
		SetResult(PluginFamilyDetails{}).
		SetPathParams(map[string]string{"id": fmt.Sprintf("%v", familyID)}).
		Get("/plugins/families/{id}")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*PluginFamilyDetails), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// Returns details for a given plugin.
//   Ref: https://cloud.tenable.com/api#/resources/plugins/plugin-details
func (c *PluginsAPI) PluginDetails(pluginID int) (*PluginDetails, error) {
	r, err := c.client.restyClient.R().
		SetResult(PluginDetails{}).
		SetPathParams(map[string]string{"id": fmt.Sprintf("%v", pluginID)}).
		Get("/plugins/plugin/{id}")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*PluginDetails), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}
