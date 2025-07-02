package instance

import (
	"fmt"
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	var dto ListRequest
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

	if dto.Tag != "" {
		queryParams["tag"] = dto.Tag
	}

	if dto.Label != "" {
		queryParams["label"] = dto.Label
	}

	if dto.MainIp != "" {
		queryParams["main_ip"] = dto.MainIp
	}

	if dto.Region != "" {
		queryParams["region"] = dto.Region
	}

	if dto.FirewallGroupId != "" {
		queryParams["firewall_group_id"] = dto.FirewallGroupId
	}

	if dto.Hostname != "" {
		queryParams["hostname"] = dto.Hostname
	}

	queryParams["show_pending_charges"] = "true"

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[interface{}](client, "https://api.vultr.com/v2/instances", queryParams, config.ApiKey)
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
			Message: "errResponse",
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

func Detail(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[interface{}](client, "https://api.vultr.com/v2/instances/"+uuid, nil, config.ApiKey)
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

func Create(c *fiber.Ctx) error {
	var dto CreateRequest
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if dto.Region == "" || dto.Plan == "" {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: "Request Body is invalid",
			Data:    nil,
		})
	}

	body := make(map[string]interface{})

	if dto.Region != "" {
		body["region"] = dto.Region
	}

	if dto.Plan != "" {
		body["plan"] = dto.Plan
	}

	if dto.OsId > 0 {
		body["os_id"] = dto.OsId
	}

	if dto.IpxeChainUrl != "" {
		body["ipxe_chain_url"] = dto.IpxeChainUrl
	}

	if dto.IsoId != "" {
		body["iso_id"] = dto.IsoId
	}

	if dto.ScriptId != "" {
		body["script_id"] = dto.ScriptId
	}

	if dto.SnapshotId != "" {
		body["snapshot_id"] = dto.SnapshotId
	}

	if dto.EnableIpv6 {
		body["enable_ipv6"] = true
	}

	if dto.DisablePublicIpv4 {
		body["disable_public_ipv4"] = true
	}

	if dto.AttachVpc != nil {
		if len(dto.AttachVpc) > 0 {
			body["attach_vpc"] = dto.AttachVpc
		}
	}

	if dto.Label != "" {
		body["label"] = dto.Label
	}

	if dto.SshKeyId != nil {
		if len(dto.SshKeyId) > 0 {
			body["sshkey_id"] = dto.SshKeyId
		}
	}

	if dto.Backups == BackupEnabled {
		body["backups"] = BackupEnabled
	} else {
		body["backups"] = BackupDisabled
	}

	if dto.AppId > 0 {
		body["app_id"] = dto.AppId
	}

	if dto.ImageId != "" {
		body["image_id"] = dto.ImageId
	}

	if dto.UserData != "" {
		body["user_data"] = dto.UserData
	}

	if dto.DdosProtection {
		body["ddos_protection"] = true
	} else {
		body["ddos_protection"] = false
	}

	if dto.ActivationEmail {
		body["activation_email"] = true
	} else {
		body["activation_email"] = false
	}

	if dto.Hostname != "" {
		body["hostname"] = dto.Hostname
	}

	if dto.FirewallGroupId != "" {
		body["firewall_group_id"] = dto.FirewallGroupId
	}

	if dto.ReversedIpv4 != "" {
		body["reserved_ipv4"] = dto.ReversedIpv4
	}

	if dto.EnableVpc {
		body["enable_vpc"] = true
	} else {
		body["enable_vpc"] = false
	}

	if dto.Tags != nil {
		if len(dto.Tags) > 0 {
			body["tags"] = dto.Tags
		}
	}

	if dto.UserScheme != "" {
		body["user_scheme"] = dto.UserScheme
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PostRequest[interface{}](client, "https://api.vultr.com/v2/instances", body, nil, config.ApiKey)
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

func Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.DeleteRequest[interface{}](client, "https://api.vultr.com/v2/instances/"+uuid, nil, nil, config.ApiKey)
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

func Start(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PostRequest[interface{}](client, "https://api.vultr.com/v2/instances/"+uuid+"/start", nil, nil, config.ApiKey)
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

func Reboot(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PostRequest[interface{}](client, "https://api.vultr.com/v2/instances/"+uuid+"/reboot", nil, nil, config.ApiKey)
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
