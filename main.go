package main

import (
	"os"

	"github.com/PauloHPMKT/goopportunities/config"
	"github.com/PauloHPMKT/goopportunities/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	port, err := initEnv()
	if err != nil {
		logger.Errorf("Error initializing env: %v", err)
		return
	}
	
	router.Initialize(port)
}

func initEnv() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	port := os.Getenv("APP_PORT")
	return port, nil
}
