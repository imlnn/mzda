package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddlewares "github.com/go-chi/chi/v5/middleware"
	"log"
	authAPI "mzda/internal/api/auth"
	subscriptionAPI "mzda/internal/api/subscription"
	userAPI "mzda/internal/api/user"
	"mzda/internal/middleware"
	"mzda/internal/storage/db/postgres"
	authSvc "mzda/internal/svc/auth"
	subSvc "mzda/internal/svc/subscription"
	userSvc "mzda/internal/svc/user"
	"net/http"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	svcName  = "AUTH"
	apiVer   = "1.0"
)

func NewAuthRouter(authService authSvc.Service) chi.Router {
	router := chi.NewRouter()

	router.Post("/signin", authAPI.SignIn(authService))
	router.Post("/renew", authAPI.RenewToken(authService))

	return router
}

func NewUserRouter(userService userSvc.Service) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.JWTAuth)

	router.Post("/changeUsername", userAPI.ChangeUsername(userService))
	router.Post("/changePassword", userAPI.ChangePassword(userService))
	router.Post("/changeEmail", userAPI.ChangeEmail(userService))

	return router
}

func NewSubscriptionRouter(subscriptionService subSvc.Service) chi.Router {
	router := chi.NewRouter()

	router.Use(chiMiddlewares.URLFormat)
	router.Use(middleware.JWTAuth)

	router.Post("/", subscriptionAPI.NewSubscription(subscriptionService))
	router.Get("/{id}", subscriptionAPI.GetSubscription(subscriptionService))
	router.Put("/{id}", subscriptionAPI.UpdateSubscription(subscriptionService))
	router.Delete("/{id}", subscriptionAPI.DeleteSubscription(subscriptionService))
	return router
}

func main() {
	// Init env
	//cfg := config.MustLoad(svcName)

	// Setup logger
	//log.Printf("Starting mzda")
	//log.Printf("Environment %v", cfg.Env)

	// TODO Setup DB
	log.Printf("Trying connect DB")
	storage, err := postgres.New()
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}
	log.Println("DB connection established")

	authService := authSvc.NewAuthSvc(storage, storage)
	userService := userSvc.NewUserSvc(storage)
	subService := subSvc.NewSubscriptionSvc(storage)

	// TODO Init server
	router := chi.NewRouter()
	root := fmt.Sprintf("/api/v%s", apiVer)

	router.Post(root+"/signup", userAPI.SignUp(userService))

	authRouter := NewAuthRouter(authService)
	router.Mount(root+"/auth", authRouter)

	userRouter := NewUserRouter(userService)
	router.Mount(root+"/user", userRouter)

	subscriptionRouter := NewSubscriptionRouter(subService)
	router.Mount(root+"/subscription", subscriptionRouter)

	err = http.ListenAndServe(":32000", router)
	if err != nil {
		return
	}
}
