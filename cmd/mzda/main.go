package main

import (
	"log"
	"mzda/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	svc      = "AUTH"
)

func main() {
	// Init env
	cfg := config.MustLoad(svc)

	// Setup logger
	log.Printf("Starting mzda")
	log.Printf("Environment %v", cfg.Env)

	// TODO Setup DB
	log.Printf("Trying connect DB")

	// TODO Init server

	// TODO Run app
}
