package main

import (
	"github.com/PauloHPMKT/goopportunities/config"
	"github.com/PauloHPMKT/goopportunities/router"
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

	port, err := config.InitializeEnv()
	if err != nil {
		logger.Errorf("Error initializing env: %v", err)
		return
	}
	
	router.Initialize(port)
}

