package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации!")
	}
	return &Config{Key: os.Getenv("KEY")}
}
