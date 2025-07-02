package plan

import (
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	config := utils.GetConfig()
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[interface{}](client, "https://api.vultr.com/v2/plans", nil, config.ApiKey)
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
