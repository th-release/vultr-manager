package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"th-release/vultr-manager/api"
	"th-release/vultr-manager/utils"
	"time"
)

func main() {
	key := flag.String("API", "", "API_KEY String *Required")
	password := flag.String("PASSWORD", "", "PASSWORD String *Required")
	port := flag.Int("PORT", 8080, "PORT Integer Default: 8080")
	addr := flag.String("DATABASE_ADDR", "localhost:5432", "DATABASE_ADDR Default: localhost:5432")
	user := flag.String("DATABASE_USER", "root", "DATABASE_USER Default: root")
	dbPassword := flag.String("DATABASE_PASSWORD", "root", "DATABASE_PASSWORD Default: root")
	schema := flag.String("DATABASE_SCHEMA", "vultrManager", "DATABASE_SCHEMA Default: vultrManager")

	flag.Parse()

	// 필수 플래그 검증
	if *key == "" {
		fmt.Println("Error: --API is required")
		flag.Usage()
		os.Exit(1)
	}
	if *password == "" {
		fmt.Println("Error: --PASSWORD is required")
		flag.Usage()
		os.Exit(1)
	}

	err := utils.SetConfig(&utils.Config{
		Port:             *port,
		ApiKey:           *key,
		Password:         *password,
		DatabaseAddr:     *addr,
		DatabaseUser:     *user,
		DatabasePassword: *dbPassword,
		DatabaseSchema:   *schema,
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
