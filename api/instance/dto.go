package instance

type ListRequest struct {
	PerPage         int64  `query:"per_page" json:"per_page"`
	Cursor          string `query:"cursor" json:"cursor"`
	Tag             string `query:"tag" json:"tag"`
	Label           string `query:"label" json:"label"`
	MainIp          string `query:"main_ip" json:"main_ip"`
	Region          string `query:"region" json:"region"`
	FirewallGroupId string `query:"firewall_group_id" json:"firewall_group_id"`
	Hostname        string `query:"hostname" json:"hostname"`
}

type DatabaseListRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type CreateRequest struct {
	Region            string     `json:"region"`
	Plan              string     `json:"plan"`
	OsId              int        `json:"os_id"`
	IpxeChainUrl      string     `json:"ipxe_chain_url"`
	IsoId             string     `json:"iso_id"`
	ScriptId          string     `json:"script_id"`
	SnapshotId        string     `json:"snapshot_id"`
	EnableIpv6        bool       `json:"enable_ipv6"`
	DisablePublicIpv4 bool       `json:"disable_public_ipv4"`
	AttachVpc         []string   `json:"attach_vpc"`
	Label             string     `json:"label"`
	SshKeyId          []string   `json:"sshkey_id"`
	Backups           BackupType `json:"backups"`
	AppId             int        `json:"app_id"`
	ImageId           string     `json:"image_id"`
	UserData          string     `json:"user_data"`
	DdosProtection    bool       `json:"ddos_protection"`
	ActivationEmail   bool       `json:"activation_email"`
	Hostname          string     `json:"hostname"`
	FirewallGroupId   string     `json:"firewall_group_id"`
	ReversedIpv4      string     `json:"reserved_ipv4"`
	EnableVpc         bool       `json:"enable_vpc"`
	Tags              []string   `json:"tags"`
	UserScheme        string     `json:"user_scheme"`
}
