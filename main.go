package main

import (
	"os"

	"github.com/PauloHPMKT/goopportunities/router"
	"github.com/joho/godotenv"
)

func main() {
	port := initEnv()
	router.InitializeRouter(port)
}

func initEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	return port
}
