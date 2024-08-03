package main

import (
	"os"

	"github.com/PauloHPMKT/goopportunities/router"
	"github.com/joho/godotenv"
)

func main() {
	port, err := initEnv()
	if err != nil {
		panic(err)
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
