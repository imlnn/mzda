package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"mzda/internal/auth/handlers"
	"mzda/internal/storage/db/postgres"
	"net/http"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	svc      = "AUTH"
	apiVer   = "1.0"
)

func main() {
	// Init env
	//cfg := config.MustLoad(svc)

	// Setup logger
	//log.Printf("Starting mzda")
	//log.Printf("Environment %v", cfg.Env)

	// TODO Setup DB
	log.Printf("Trying connect DB")
	storage, err := postgres.New()
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}

	// TODO Init server
	router := chi.NewRouter()
	router.Post("/signin", handlers.SignIn(storage, storage))
	router.Post("/signup", handlers.SignUp(storage))

	err := http.ListenAndServe(":32000", router)
	if err != nil {
		return
	}
}
