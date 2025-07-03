package region

import "th-release/vultr-manager/utils"

type RegionListResponse struct {
	Regions []Region   `json:"regions"`
	Meta    utils.Meta `json:"meta"`
}

type Region struct {
	ID        string   `json:"id"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Continent string   `json:"continent"`
	Options   []string `json:"options"`
}
