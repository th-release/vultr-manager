package main

import (
	"fmt"
	"log"
	"th-release/vultr-manager/api"
	"th-release/vultr-manager/utils"
	"time"
)

func main() {
	config := utils.GetConfig()

	if config == nil {
		log.Fatalln("Not Found Config")
		return
	}

	app := api.InitServer(config)

	if app == nil {
		log.Fatalln("Init Server Error")
		return
	}

	log.Println(time.Now())

	app.App.Listen(fmt.Sprintf(":%s", config.Port))
}
