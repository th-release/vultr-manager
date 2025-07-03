package application

import (
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
)

func ApplicationList(queryParams map[string]string, token string) (*resty.Response, ListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[ListResponse](client, "https://api.vultr.com/v2/applications", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}
