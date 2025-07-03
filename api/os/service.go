package os

import (
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
)

func OsList(queryParams map[string]string, token string) (*resty.Response, OsListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[OsListResponse](client, "https://api.vultr.com/v2/os", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}
