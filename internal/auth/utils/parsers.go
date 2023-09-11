package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"mzda/internal/auth/handlers"
	"mzda/internal/storage/models/mzda"
)

func ParseUserDTO(b io.ReadCloser) (*mzda.UserDTO, error) {
	const fn = "internal/auth/utils/users/ParseUser"
	var usr mzda.UserDTO

	err := json.NewDecoder(b).Decode(&usr)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &usr, nil
}

func ParseChangePassword(b io.ReadCloser) (*handlers.ChangePasswordRequest, error) {
	const fn = "internal/auth/utils/users/ParseChangePassword"
	var req handlers.ChangePasswordRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func ParseChangeUsername(b io.ReadCloser) (*handlers.ChangeUsernameRequest, error) {
	const fn = "internal/auth/utils/users/ParseChangePassword"
	var req handlers.ChangeUsernameRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func ParseChangeEmail(b io.ReadCloser) (*handlers.ChangeEmailRequest, error) {
	const fn = "internal/auth/utils/users/ParseChangePassword"
	var req handlers.ChangeEmailRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func ParseCredentials(b io.ReadCloser) (*handlers.Credentials, error) {
	const fn = "internal/auth/utils/users/ParseCredentials"
	var req handlers.Credentials

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s  %v", fn, err)
	}

	return &req, nil
}
