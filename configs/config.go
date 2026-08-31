package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DbConfig
	Auth AuthConfig
	SMTP SMTPConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type SMTPConfig struct {
	Host string
	Password string
	Email string
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	config := Config{
		DB: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("Secret"),
		},
		SMTP: SMTPConfig{
			Host: os.Getenv("SMTP_HOST"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Email: os.Getenv("SMTP_EMAIL"),
			Port: os.Getenv("SMTP_PORT"),
		},
	}

	return &config
}