package script

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

	config := utils.GetConfig()
	resp, res, errResp, err := ScriptList(queryParams, config.ApiKey)
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

func Detail(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[DetailResponse](client, "https://api.vultr.com/v2/startup-scripts/"+uuid, nil, config.ApiKey)
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

	decode, err := utils.DecodeBase64(res.StartUpScript.Script)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	res.StartUpScript.Script = decode

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    res,
	})
}

func Create(c *fiber.Ctx) error {
	var dto ScriptCreateRequest
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if dto.Name == "" || dto.Script == "" || !dto.Type.IsValid() {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: "Request Body is invalid",
			Data:    nil,
		})
	}

	body := map[string]interface{}{}

	if dto.Name != "" {
		body["name"] = dto.Name
	}

	if dto.Script != "" {
		body["script"] = utils.EncodeBase64(dto.Script)
	}

	if dto.Type.String() != "" {
		body["type"] = dto.Type.String()
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PostRequest[map[string]interface{}](client, "https://api.vultr.com/v2/startup-scripts", body, nil, config.ApiKey)
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

func Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	var dto ScriptUpdateRequest
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	queryParams := map[string]interface{}{}

	if dto.Name != "" {
		queryParams["name"] = dto.Name
	}

	if dto.Script != "" {
		queryParams["script"] = utils.EncodeBase64(dto.Script)
	}

	if dto.Type.String() != "" {
		queryParams["type"] = dto.Type.String()
	}

	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.PatchRequest[interface{}](client, "https://api.vultr.com/v2/startup-scripts/"+uuid, queryParams, nil, config.ApiKey)
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
	resp, res, errResp, err := utils.DeleteRequest[map[string]interface{}](client, "https://api.vultr.com/v2/startup-scripts/"+uuid, nil, nil, config.ApiKey)
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
