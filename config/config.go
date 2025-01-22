package config

import (
	"log"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
	DBDSN  string
	SYMBOLS []string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		APIKey: os.Getenv("API_KEY"),
		DBDSN:  os.Getenv("DB_DSN"),
		SYMBOLS: strings.Split(os.Getenv("SYMBOLS"), ","),
	}
}
