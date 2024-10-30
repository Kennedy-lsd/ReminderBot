package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TELEGRAM_APITOKEN      string
	PORT                   string
	POSTGRES_HOST          string
	POSTGRES_USER          string
	POSTGRES_PASSWORD      string
	POSTGRES_DATABASE_NAME string
	POSTGRES_PORT          string
	SSL_MODE               string
}

func ConfigInit() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return &Config{
		TELEGRAM_APITOKEN:      os.Getenv("TELEGRAM_APITOKEN"),
		PORT:                   os.Getenv("PORT"),
		POSTGRES_HOST:          os.Getenv("POSTGRES_HOST"),
		POSTGRES_USER:          os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD:      os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DATABASE_NAME: os.Getenv("POSTGRES_DATABASE_NAME"),
		POSTGRES_PORT:          os.Getenv("POSTGRES_PORT"),
		SSL_MODE:               os.Getenv("SSL_MODE"),
	}

}
