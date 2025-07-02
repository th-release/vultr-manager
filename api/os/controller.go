package os

import (
	"fmt"
	"th-release/vultr-manager/utils"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	var dto ListRequest
	if err := c.BodyParser(&dto); err != nil {
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

	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[interface{}](client, "https://api.vultr.com/v2/os", queryParams, "")
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
