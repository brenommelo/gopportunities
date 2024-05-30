package main

import (
	"github.com/brenommelo/gopportunities/config"
	"github.com/brenommelo/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Confir initialization error: %v", err)
		return
	}

	//initialize router
	router.Initialize()
}
