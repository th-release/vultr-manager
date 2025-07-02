package firewall

type FireWallGroupListRequest struct {
	PerPage int64  `query:"per_page" json:"per_page"`
	Cursor  string `query:"cursor" json:"cursor"`
}

type FireWallRulesCreateRequest struct {
	IpType     IpType   `json:"ip_type"`
	Protocol   Protocol `json:"protocol"`
	Subnet     string   `json:"subnet"`
	SubnetSize int      `json:"subnet_size"`
	Port       string   `json:"port"`
	Source     string   `json:"source"`
	Notes      string   `json:"notes"`
}

type CreateFirewallGroupRequest struct {
	Description string `json:"description"`
}

type UpdateFirewallGroupRequest struct {
	Description string `json:"description"`
}

type FireWallRulesListRequest struct {
	PerPage int64  `query:"per_page" json:"per_page"`
	Cursor  string `query:"cursor" json:"cursor"`
}
