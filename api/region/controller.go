package region

import (
	"fmt"
	"th-release/vultr-manager/utils"

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

	resp, res, errResp, err := RegionList(queryParams, "")
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
