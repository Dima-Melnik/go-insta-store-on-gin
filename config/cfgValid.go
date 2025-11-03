package config

import (
	"log"

	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/utils"
)

func (c *Config) ValidateConfig() {
	if c.Database.DBPassword == "" {
		dbPassword, ok := utils.GetEnv("DB_PASSWORD")
		if !ok {
			return
		}
		c.Database.DBPassword = dbPassword
	}

	if c.Database.DBPort == 0 {
		c.Database.DBPort = 5432
	}

	if c.Server.Port == 0 {
		c.Server.Port = 8080
	}

	if c.Database.DBSSLMode == "" {
		c.Database.DBSSLMode = "disable"
	}

	if c.Database.DBHost == "" || c.Database.DBUser == "" || c.Database.DBName == "" {
		log.Println("error: fields can not be empty")
	}
}
