package application

import "th-release/vultr-manager/utils"

type ListResponse struct {
	Applications []Application `json:"applications"`
	Meta         utils.Meta    `json:"meta"`
}

type Application struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	DeployName string `json:"deploy_name"`
	Type       string `json:"type"`
	Vendor     string `json:"vendor"`
	ImageId    string `json:"image_id"`
}
