package utils

import (
	"os"
	"strconv"
)

func GetConfig() *Config {
	boolean, err := strconv.ParseBool(os.Getenv("SYNC_DATABASE"))

	if err != nil {
		boolean = false
	}

	return &Config{
		ApiKey:           ThreeTermString(len(os.Getenv("API")) >= 0, os.Getenv("API"), "APIKEY"),
		Password:         ThreeTermString(len(os.Getenv("PASSWORD")) >= 0, os.Getenv("PASSWORD"), "PASSWORD"),
		Port:             ThreeTermString(len(os.Getenv("PORT")) >= 0, os.Getenv("PORT"), "8080"),
		DatabaseAddr:     ThreeTermString(len(os.Getenv("DATABASE_ADDR")) >= 0, os.Getenv("DATABASE_ADDR"), "localhost"),
		DatabaseUser:     ThreeTermString(len(os.Getenv("DATABASE_USER")) >= 0, os.Getenv("DATABASE_USER"), "vultr"),
		DatabasePassword: ThreeTermString(len(os.Getenv("DATABASE_PASSWORD")) >= 0, os.Getenv("DATABASE_PASSWORD"), "vultr"),
		DatabaseSchema:   ThreeTermString(len(os.Getenv("DATABASE_SCHEMA")) >= 0, os.Getenv("DATABASE_SCHEMA"), "vultr"),
		SyncDatabase:     boolean,
	}
}
