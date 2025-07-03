package plan

import "th-release/vultr-manager/utils"

type PlanListResponse struct {
	Plans []Plan     `json:"plans"`
	Meta  utils.Meta `json:"meta"`
}

type Plan struct {
	Bandwidth   int      `json:"bandwidth"`
	CpuVendor   string   `json:"cpu_vendor"`
	Disk        int      `json:"disk"`
	DiskCount   int      `json:"disk_count"`
	DiskType    string   `json:"disk_type"`
	HourlyCost  float64  `json:"hourly_cost"`
	ID          string   `json:"id"`
	InvoiceType string   `json:"invoice_type"`
	Locations   []string `json:"locations"`
	MonthlyCost float64  `json:"monthly_cost"`
	RAM         int      `json:"ram"`
	Type        string   `json:"type"`
	VcpuCount   int      `json:"vcpu_count"`
}
