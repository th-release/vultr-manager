package firewall

import (
	"fmt"
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func FireWallGroupList(c *fiber.Ctx) error {
	var dto FireWallGroupListRequest
	if err := c.QueryParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if dto.PerPage == 0 {
		dto.PerPage = 100
	}

	queryParams := map[string]string{
		"per_page": fmt.Sprintf("%d", dto.PerPage),
	}
	if dto.Cursor != "" {
		queryParams["cursor"] = dto.Cursor
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[FirewallGroupListResponse](client, "https://api.vultr.com/v2/firewalls", queryParams, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallGroupDetail(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[FirewallGroupDetailResponse](client, "https://api.vultr.com/v2/firewalls/"+uuid, nil, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallGroupCreate(c *fiber.Ctx) error {
	var dto CreateFirewallGroupRequest
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	body := map[string]string{}

	if dto.Description != "" {
		body["description"] = dto.Description
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PostRequest[FirewallGroupDetailResponse](client, "https://api.vultr.com/v2/firewalls", body, nil, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(201).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallGroupUpdate(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	var dto UpdateFirewallGroupRequest
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	body := map[string]interface{}{}

	if dto.Description != "" {
		body["description"] = dto.Description
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PutRequest[interface{}](client, "https://api.vultr.com/v2/firewalls/"+uuid, body, nil, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallGroupDelete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.DeleteRequest[interface{}](client, "https://api.vultr.com/v2/firewalls/"+uuid, nil, nil, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallRulesList(c *fiber.Ctx) error {
	group := c.Params("group")

	var dto FireWallRulesListRequest
	if err := c.QueryParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if dto.PerPage == 0 {
		dto.PerPage = 100
	}

	queryParams := map[string]string{
		"per_page": fmt.Sprintf("%d", dto.PerPage),
	}
	if dto.Cursor != "" {
		queryParams["cursor"] = dto.Cursor
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[FireWallRulesListResponse](client, "https://api.vultr.com/v2/firewalls/"+group+"/rules", queryParams, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallRulesDetail(c *fiber.Ctx) error {
	group := c.Params("group")
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[interface{}](client, "https://api.vultr.com/v2/firewalls/"+group+"/rules/"+uuid, nil, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallRulesCreate(c *fiber.Ctx) error {
	group := c.Params("group")

	var dto FireWallRulesCreateRequest
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	body := make(map[string]interface{})

	if dto.IpType.IsValid() {
		body["ip_type"] = dto.IpType.String()
	}

	if dto.Protocol.IsValid() {
		body["protocol"] = dto.Protocol.String()
	}

	if dto.Subnet != "" {
		body["subnet"] = dto.Subnet
	}

	body["subnet_size"] = dto.SubnetSize

	if !dto.IpType.IsValid() || !dto.Protocol.IsValid() || dto.Subnet == "" || dto.SubnetSize < 0 {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: "Request Body is invalid",
			Data:    nil,
		})
	}

	if dto.Protocol != ProtocolTCP && dto.Protocol != ProtocolUDP && dto.Port != "" {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: "Request Body is invalid",
			Data:    nil,
		})
	}

	if dto.Port != "" {
		body["port"] = dto.Port
	}

	if dto.Source != "" {
		body["source"] = dto.Source
	}

	if dto.Notes != "" {
		body["notes"] = dto.Notes
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PostRequest[interface{}](client, "https://api.vultr.com/v2/firewalls/"+group+"/rules", body, nil, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(201).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func FireWallRulesDelete(c *fiber.Ctx) error {
	group := c.Params("group")
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.DeleteRequest[interface{}](client, "https://api.vultr.com/v2/firewalls/"+group+"/rules/"+uuid, nil, nil, config.ApiKey)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if errResp != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "",
			Data:    errResp,
		})
	}

	if err := utils.CheckResponse(resp); err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}
