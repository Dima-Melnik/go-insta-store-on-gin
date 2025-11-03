package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		DBHost     string `yaml:"db_host"`
		DBPort     int    `yaml:"db_port"`
		DBUser     string `yaml:"db_user"`
		DBPassword string `yaml:"db_password"`
		DBName     string `yaml:"db_name"`
		DBSSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

var Cfg Config

func LoadConfig(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(f, &Cfg); err != nil {
		log.Printf("failed unmarshaling yaml %s: %v", path, err)
		return err
	}

	Cfg.ValidateConfig()

	return nil
}

func (c *Config) InitServerPort() string {
	port := fmt.Sprintf(":%d", c.Server.Port)

	return port
}

func (c *Config) InitDatabaseDsn() string {
	dsn := fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=%s",
		c.Database.DBHost,
		c.Database.DBPort,
		c.Database.DBPassword,
		c.Database.DBUser,
		c.Database.DBName,
		c.Database.DBSSLMode)

	return dsn
}
