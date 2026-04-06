package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB DbConfig
}

type DbConfig struct {
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		DB: DbConfig{
			DSN: os.Getenv("DSN"),
		},
	}
}
