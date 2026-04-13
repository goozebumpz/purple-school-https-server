package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DbConfig struct {
	DSN string
}

type AuthConfig struct {
	Secret string
}

type Config struct {
	Db         DbConfig
	AuthConfig AuthConfig
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Db: DbConfig{
			DSN: os.Getenv("DSN"),
		},
		AuthConfig: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
