package firewall

import "th-release/vultr-manager/utils"

type FirewallGroupListResponse struct {
	Firewalls []FirewallGroup `json:"firewall_groups"`
	Meta      utils.Meta      `json:"meta"`
}

type FirewallGroupDetailResponse struct {
	FirewallGroup FirewallGroup `json:"firewall_group"`
}

type FirewallGroup struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	DateCreated   string `json:"date_created"`
	DateModified  string `json:"date_modified"`
	InstanceCount int    `json:"instance_count"`
	RuleCount     int    `json:"rule_count"`
	MaxRuleCount  int    `json:"max_rule_count"`
}

type FireWallRulesListResponse struct {
	FirewallRules []FirewallRule `json:"firewall_rules"`
	Meta          utils.Meta     `json:"meta"`
}

type FirewallRule struct {
	Id         int64  `json:"id"`
	Type       string `json:"type"`
	IpType     string `json:"ip_type"`
	Action     string `json:"action"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
	Subnet     string `json:"subnet"`
	SubnetSize int64  `json:"subnet_size"`
	Source     string `json:"source"`
	Notes      string `json:"notes"`
}
