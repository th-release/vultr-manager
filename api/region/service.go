package region

import (
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
)

func RegionList(queryParams map[string]string, token string) (*resty.Response, RegionListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[RegionListResponse](client, "https://api.vultr.com/v2/regions", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}
