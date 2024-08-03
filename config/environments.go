package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnv() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	port := os.Getenv("APP_PORT")
	return port, nil
}
