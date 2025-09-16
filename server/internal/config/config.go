package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error while loading .env file: ", err)
	}

	port := os.Getenv("PORT")
	c.Port = port
}
