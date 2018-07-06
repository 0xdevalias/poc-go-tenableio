package api

import "time"

type Filter struct {
	Filter  string `json:"filter"`
	Quality string `json:"quality"`
	Value   string `json:"value"`
}

// Ref: https://cloud.tenable.com/api#/resources/plugins
type PluginFamilies struct {
	Families []PluginFamily `json:"families"`
}

// Ref: https://cloud.tenable.com/api#/resources/plugins
type PluginFamily struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Ref: https://cloud.tenable.com/api#/resources/plugins/family-details
type PluginFamilyDetails struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Plugins []Plugin `json:"plugins"`
}

// Ref: https://cloud.tenable.com/api#/resources/plugins
type Plugin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Ref: https://cloud.tenable.com/api#/resources/plugins/plugin-details
type PluginDetails struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	FamilyName string             `json:"family_name"`
	Attributes []PluginAttributes `json:"attributes"`
}

// Ref: https://cloud.tenable.com/api#/resources/plugins
type PluginAttributes struct {
	Name  string `json:"attribute_name"`
	Value string `json:"attribute_value"`
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
	Description              string                         `json:"description"`
	Discovery                interface{}                    `json:"discovery"`
	PluginDetails            VulnerabilityInfoPluginDetails `json:"plugin_details"`
	ReferenceInformation     interface{}                    `json:"reference_information"`
	RiskInformation          interface{}                    `json:"risk_information"`
	SeeAlso                  []interface{}                  `json:"see_also"`
	Solution                 string                         `json:"solution"`
	VulnerabilityInformation VulnerabilityInformation       `json:"vulnerability_information"`
}

type VulnerabilityInfoPluginDetails struct {
	Name             string     `json:"name"`
	FamilyName       string     `json:"family"`
	Severity         string     `json:"severity"`
	Type             string     `json:"local"`
	PublicationDate  *time.Time `json:"publication_date"`
	ModificationDate *time.Time `json:"modification_date"`
	Version          string     `json:"version"`
}

// vulnerability_info
//   Ref: https://cloud.tenable.com/api#/resources/workbenches
type VulnerabilityInformation struct {
	CPE                  []string   `json:"cpe"`
	VulnPublicationDate  *time.Time `json:"vulnerability_publication_date"`
	PatchPublicationDate *time.Time `json:"patch_publication_date"`
	UnsupportedByVendor  bool       `json:"unsupported_by_vendor"`

	// TODO: Finish implementing this..
	ExploitAvailable   interface{} `json:"exploit_available"`
	ExploitabilityEase interface{} `json:"exploitability_ease"`
	ExploitedByMalware interface{} `json:"exploited_by_malware"`
	ExploitedByNessus  interface{} `json:"exploited_by_nessus"`
	ExploitFrameworks  interface{} `json:"exploit_frameworks"`
	AssetInventory     interface{} `json:"asset_inventory"`
	DefaultAccount     interface{} `json:"default_account"`
	InTheNews          interface{} `json:"in_the_news"`
	Malware            interface{} `json:"malware"`
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
