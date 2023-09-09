package main

import (
	"fmt"
	"log"
	"mzda/internal/storage/db/postgres"
	"mzda/internal/storage/models/mzda"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	svc      = "AUTH"
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

	usrDTO1 := mzda.UserDTO{
		Username: "TestUser1",
		Pwd:      "1234",
		Email:    "test1@test.com",
	}

	usrDTO2 := mzda.UserDTO{
		Username: "TestUser2",
		Pwd:      "4321",
		Email:    "test2@test.com",
	}

	err = storage.AddUser(&usrDTO1)
	if err != nil {
		fmt.Println(err)
	}

	err = storage.AddUser(&usrDTO2)
	if err != nil {
		fmt.Println(err)
	}

	usr, err := storage.UserByName("TestUser1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr)

	usrID := usr.ID

	usr, err = storage.UserByEmail("test2@test.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr)

	usr, err = storage.UserByID(usrID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr)

	err = storage.DeleteUser(usr)
	if err != nil {
		fmt.Println(err)
	}

	usr, err = storage.UserByEmail("test2@test.com")
	usr.Username = "TestUser"

	err = storage.UpdateUser(usr)
	if err != nil {
		fmt.Println(err)
	}

	usr, err = storage.UserByName("TestUser")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usr)

	// TODO Init server

	// TODO Run app
}
