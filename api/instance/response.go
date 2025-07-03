package instance

import "th-release/vultr-manager/utils"

type InstsanceListResponse struct {
	Instances []Instance `json:"instances"`
	Meta      utils.Meta `json:"meta"`
}

type Instance struct {
	AllowedBandwidth int      `json:"allowed_bandwidth"`
	AppID            int      `json:"app_id"`
	DateCreated      string   `json:"date_created"`
	Disk             int      `json:"disk"`
	Features         []string `json:"features"`
	FirewallGroupID  string   `json:"firewall_group_id"`
	GatewayV4        string   `json:"gateway_v4"`
	Hostname         string   `json:"hostname"`
	ID               string   `json:"id"`
	ImageID          string   `json:"image_id"`
	InternalIP       string   `json:"internal_ip"`
	KVM              string   `json:"kvm"`
	Label            string   `json:"label"`
	MainIP           string   `json:"main_ip"`
	NetmaskV4        string   `json:"netmask_v4"`
	OS               string   `json:"os"`
	OsID             int      `json:"os_id"`
	PendingCharges   float64  `json:"pending_charges"`
	Plan             string   `json:"plan"`
	PowerStatus      string   `json:"power_status"`
	RAM              int      `json:"ram"`
	Region           string   `json:"region"`
	ServerStatus     string   `json:"server_status"`
	Status           string   `json:"status"`
	Tag              string   `json:"tag"`
	Tags             []string `json:"tags"`
	UserScheme       string   `json:"user_scheme"`
	V6MainIP         string   `json:"v6_main_ip"`
	V6Network        string   `json:"v6_network"`
	V6NetworkSize    int      `json:"v6_network_size"`
	VcpuCount        int      `json:"vcpu_count"`
	Vpcs             []string `json:"vpcs"`
}
