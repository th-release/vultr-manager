package api

import (
	"th-release/vultr-manager/utils"

	"github.com/gofiber/fiber/v2"
)

func ApiMiddleware(c *fiber.Ctx) error {
	var dto GeneralRequest

	method := c.Method()

	if method == "POST" || method == "PUT" || method == "DELETE" || method == "PATCH" {
		if err := c.BodyParser(&dto); err != nil {
			return c.Status(400).JSON(utils.BasicResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
	} else {
		if err := c.QueryParser(&dto); err != nil {
			return c.Status(400).JSON(utils.BasicResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}

	config := utils.GetConfig()

	if dto.Password != config.Password {
		return c.Status(401).JSON(utils.BasicResponse{
			Success: false,
			Message: "Invalid Password",
			Data:    nil,
		})
	}

	return c.Next()
}

func ApplicationMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func RegionMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func PlanMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func FirewallMiddleware(c *fiber.Ctx) error {
	config := utils.GetConfig()
	if config == nil {
		return c.Status(401).JSON(utils.BasicResponse{
			Success: false,
			Message: "Unauthorized",
			Data:    nil,
		})
	}

	return c.Next()
}

func InstanceMiddleware(c *fiber.Ctx) error {
	config := utils.GetConfig()
	if config == nil {
		return c.Status(401).JSON(utils.BasicResponse{
			Success: false,
			Message: "Unauthorized",
			Data:    nil,
		})
	}

	return c.Next()
}

func ScriptMiddleware(c *fiber.Ctx) error {
	config := utils.GetConfig()
	if config == nil {
		return c.Status(401).JSON(utils.BasicResponse{
			Success: false,
			Message: "Unauthorized",
			Data:    nil,
		})
	}
	return c.Next()
}

func OsMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
