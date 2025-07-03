package utils

type BasicResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Config struct {
	ApiKey           string `json:"APIKEY"`
	Password         string `json:"PASSWORD"`
	Port             int    `json:"PORT"`
	DatabaseAddr     string `json:"DATABASE_ADDR"`
	DatabaseUser     string `json:"DATABASE_USER"`
	DatabasePassword string `json:"DATABASE_PASSWORD"`
	DatabaseSchema   string `json:"DATABASE_SCHEMA"`
}

type Meta struct {
	Total int64 `json:"total"`
	Links Links `json:"links"`
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
