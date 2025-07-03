package plan

import (
	"th-release/vultr-manager/utils"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	config := utils.GetConfig()
	resp, res, errResp, err := PlanList(nil, config.ApiKey)
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
