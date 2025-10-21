package utils

import (
	"os"
)

func GetConfig() *Config {
	return &Config{
		ApiKey:           ThreeTermString(len(os.Getenv("API")) >= 0, os.Getenv("API"), "APIKEY"),
		Password:         ThreeTermString(len(os.Getenv("PASSWORD")) >= 0, os.Getenv("API"), "PASSWORD"),
		Port:             ThreeTermString(len(os.Getenv("PORT")) >= 0, os.Getenv("PORT"), "3000"),
		DatabaseAddr:     ThreeTermString(len(os.Getenv("DATABASE_ADDR")) >= 0, os.Getenv("DATABASE_ADDR"), "localhost"),
		DatabaseUser:     ThreeTermString(len(os.Getenv("DATABASE_USER")) >= 0, os.Getenv("DATABASE_USER"), "vultr"),
		DatabasePassword: ThreeTermString(len(os.Getenv("DATABASE_PASSWORD")) >= 0, os.Getenv("DATABASE_PASSWORD"), "vultr"),
		DatabaseSchema:   ThreeTermString(len(os.Getenv("DATABASE_SCHEMA")) >= 0, os.Getenv("DATABASE_SCHEMA"), "vultr"),
	}
}
