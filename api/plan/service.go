package plan

import (
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
)

func PlanList(queryParams map[string]string, token string) (*resty.Response, PlanListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[PlanListResponse](client, "https://api.vultr.com/v2/plans", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}
