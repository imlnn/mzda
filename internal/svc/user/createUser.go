package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mzda/internal/storage/models"
	"net/http"
)

func parseUserDTO(b io.ReadCloser) (*models.UserDTO, error) {
	const fn = "internal/svc/user/createUser/parseUserDTO"
	var usr models.UserDTO

	err := json.NewDecoder(b).Decode(&usr)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &usr, nil
}

func (svc *UserSvc) CreateUser(req *http.Request) (err error, statusCode int) {
	const fn = "internal/svc/user/createUser/CreateUser"
	usr, err := parseUserDTO(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to parse request"), http.StatusBadRequest
	}
	err = svc.userStorage.AddUser(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to register"), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}
