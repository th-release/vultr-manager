package instance

import (
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
)

func InstanceList(queryParams map[string]string, token string) (*resty.Response, InstsanceListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[InstsanceListResponse](client, "https://api.vultr.com/v2/instances", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}
