package main

import (
	"flag"
	"fmt"
	"log"
	"th-release/vultr-manager/api"
	"th-release/vultr-manager/utils"
	"time"
)

func main() {
	key := flag.String("API", "", "API_KEY")
	password := flag.String("PASSWORD", "", "PASSWORD")
	port := flag.Int("PORT", 8080, "PORT")

	flag.Parse()

	err := utils.SetConfig(&utils.Config{
		Port:     *port,
		ApiKey:   *key,
		Password: *password,
	})

	if err != nil {
		log.Fatalf("Config 저장중 에러 발생 %s", err.Error())
	}

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

	app.App.Listen(fmt.Sprintf(":%d", config.Port))
}
