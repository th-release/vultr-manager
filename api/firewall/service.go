package firewall

import (
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
)

func FirewallGroupList(queryParams map[string]string, token string) (*resty.Response, FirewallGroupListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[FirewallGroupListResponse](client, "https://api.vultr.com/v2/firewalls", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}

func FirewallRulesList(group string, queryParams map[string]string, token string) (*resty.Response, FireWallRulesListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[FireWallRulesListResponse](client, "https://api.vultr.com/v2/firewalls/"+group+"/rules", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}
