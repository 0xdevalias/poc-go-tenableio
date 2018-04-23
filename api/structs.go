package api

type Filter struct {
	Filter  string `json:"filter"`
	Quality string `json:"quality"`
	Value   string `json:"value"`
}

// Ref: https://cloud.tenable.com/api#/resources/workbenches/assets
type Assets struct {
	Assets []Asset `json:"assets"`
}

// Ref: https://cloud.tenable.com/api#/resources/workbenches/assets
type Asset struct {
	ID              string   `json:"id""`
	HasAgent        bool     `json:"has_agent"`
	LastSeen        string   `json:"last_seen"`
	Sources         []Source `json:"sources"`
	AgentName       []string `json:"agent_name"`
	IPv4            []string `json:"ipv4"`
	IPv6            []string `json:"ipv6"`
	FQDN            []string `json:"fqdn"`
	NetbiosName     []string `json:"netbios_name"`
	OperatingSystem []string `json:"operating_system"`
	MacAddress      []string `json:"mac_address"`
}

// Ref: https://cloud.tenable.com/api#/resources/workbenches/asset-vulnerabilities
type AssetVulnerabilities struct {
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

// Ref: https://cloud.tenable.com/api#/resources/workbenches/asset-vulnerability-info
type AssetVulnInfo struct {
	Info VulnerabilityInfo `json:"info"`
}

// Ref: https://cloud.tenable.com/api#/resources/workbenches/asset-vulnerability-output
type AssetVulnOutputs struct {
	Outputs []AssetVulnOutput `json:"outputs"`
}

// Ref: https://cloud.tenable.com/api#/resources/workbenches/asset-vulnerability-output
// TODO: Finish the AssetVulnOutput struct
type AssetVulnOutput struct {
	PluginOutput string      `json:"plugin_output"`
	States       interface{} `json:"states"`
}

// vulnerability
//   Ref: https://cloud.tenable.com/api#/resources/workbenches
type Vulnerability struct {
	Count              int    `json:"count"`
	PluginFamily       string `json:"plugin_family"`
	PluginID           int    `json:"plugin_id"`
	PluginName         string `json:"plugin_name"`
	VulnerabilityState string `json:"vulnerability_state"`
	Severity           int    `json:"severity"`
}

// vulnerability_info
//   Ref: https://cloud.tenable.com/api#/resources/workbenches
// TODO: Finish the VulnerabilityInfo struct
type VulnerabilityInfo struct {
	Description              string                   `json:"description"`
	Discovery                interface{}              `json:"discovery"`
	PluginDetails            interface{}              `json:"plugin_details"`
	ReferenceInformation     interface{}              `json:"reference_information"`
	RiskInformation          interface{}              `json:"risk_information"`
	SeeAlso                  []interface{}            `json:"see_also"`
	Solution                 string                   `json:"solution"`
	VulnerabilityInformation VulnerabilityInformation `json:"vulnerability_information"`
}

// TODO: Finish implementing this..
type VulnerabilityInformation struct {
	CPE []string `json:"cpe"`
}

// vulnerability_output
//   Ref: https://cloud.tenable.com/api#/resources/workbenches
type VulnerabilityOutput struct {
	ApplicationProtocol string  `json:"application_protocol"`
	Assets              []Asset `json:"assets"`
	Port                int     `json:"port"`
	TransportProtocol   string  `json:"transport_protocol"`
}

// source
//   Ref: https://cloud.tenable.com/api#/resources/workbenches
type Source struct {
	Name      string `json:"name"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}
