package api

import (
	"th-release/vultr-manager/api/application"
	"th-release/vultr-manager/api/firewall"
	"th-release/vultr-manager/api/instance"
	"th-release/vultr-manager/api/os"
	"th-release/vultr-manager/api/plan"
	"th-release/vultr-manager/api/region"
	"th-release/vultr-manager/api/script"
	"th-release/vultr-manager/utils"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	App    *fiber.App
	Config utils.Config
}

func InitServer(config *utils.Config) *ServerConfig {
	app := fiber.New()

	if config == nil {
		return nil
	}

	server := &ServerConfig{
		App:    app,
		Config: *config,
	}

	server.setupRoutes()
	return server
}

func (s *ServerConfig) setupRoutes() {
	apiGroup := s.App.Group("/api", ApiMiddleware)

	apiGroup.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(utils.BasicResponse{
			Success: true,
			Message: "",
			Data:    nil,
		})
	})

	osGroup := apiGroup.Group("/os", OsMiddleware)
	osGroup.Get("/list", os.List)

	appicationGroup := apiGroup.Group("/application", ApplicationMiddleware)
	appicationGroup.Get("/list", application.List)

	regionGroup := apiGroup.Group("/region", RegionMiddleware)
	regionGroup.Get("/list", region.List)

	planGroup := apiGroup.Group("/plan", PlanMiddleware)
	planGroup.Get("/list", plan.List)

	firewallGroup := apiGroup.Group("/firewall", FirewallMiddleware)

	firewallGroup.Get("/group/list", firewall.FireWallGroupList)
	firewallGroup.Get("/group/detail/:uuid", firewall.FireWallGroupDetail)
	firewallGroup.Post("/group/create", firewall.FireWallGroupCreate)
	firewallGroup.Put("/group/update/:uuid", firewall.FireWallGroupUpdate)
	firewallGroup.Delete("/group/delete/:uuid", firewall.FireWallGroupDelete)

	firewallGroup.Get("/rules/:group/list", firewall.FireWallRulesList)
	firewallGroup.Get("/rules/:group/detail/:uuid", firewall.FireWallRulesDetail)
	firewallGroup.Post("/rules/:group/create", firewall.FireWallRulesCreate)
	firewallGroup.Delete("/rules/:group/delete/:uuid", firewall.FireWallRulesDelete)

	scriptGroup := apiGroup.Group("/script", ScriptMiddleware)
	scriptGroup.Get("/list", script.List)
	scriptGroup.Get("/detail/:uuid", script.Detail)
	scriptGroup.Post("/create", script.Create)
	scriptGroup.Put("/update/:uuid", script.Update)
	scriptGroup.Delete("/delete/:uuid", script.Delete)

	instanceGroup := apiGroup.Group("/instance", InstanceMiddleware)
	instanceGroup.Get("/list", instance.List)
	instanceGroup.Get("/detail/:uuid", instance.Detail)
	instanceGroup.Post("/create", instance.Create)
	instanceGroup.Delete("/delete/:uuid", instance.Delete)
	instanceGroup.Post("/start/:uuid", instance.Start)
}
