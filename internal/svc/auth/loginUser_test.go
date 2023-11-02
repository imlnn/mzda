package auth

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"mzda/internal/storage/db/mock"
	"mzda/internal/storage/models"
	"net/http"
	"strings"
	"testing"
)

func TestLoginUser_Successful(t *testing.T) {
	// Create a mock AuthsStorage
	storage := mock.NewMockConnection()

	// Create Auth service
	authSvc := NewAuthSvc(storage, storage)

	// Create a sample user
	userDTO := &models.UserDTO{
		Username: "testuser",
		Pwd:      "password123",
		Email:    "sample@acme.com",
	}

	_ = authSvc.userStorage.AddUser(userDTO)
	sampleUser, _ := authSvc.userStorage.UserByName(userDTO.Username)

	// Create a JSON payload for the request
	loginReq := loginRequest{
		Username: sampleUser.Username,
		Password: sampleUser.Pwd,
	}
	reqPayload, _ := json.Marshal(loginReq)

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(reqPayload))
	assert.NoError(t, err)

	// Call the LoginUser function
	response, statusCode, err := authSvc.LoginUser(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Decode the response
	var loginRes loginResponse
	err = json.Unmarshal(response, &loginRes)
	assert.NoError(t, err)

	// Check if the JWT and Refresh token were generated
	assert.NotEqual(t, "", loginRes.JWT)
	assert.NotEqual(t, "", loginRes.Refresh)

	auth, err := storage.GetAuth(loginRes.Refresh)
	assert.NoError(t, err)
	assert.Equal(t, auth.RefreshToken, loginRes.Refresh)
}

func TestLoginUser_InvalidUser(t *testing.T) {
	// Create a mock AuthsStorage
	storage := mock.NewMockConnection()

	// Create a JSON payload for the request with invalid user credentials
	loginReq := loginRequest{
		Username: "invaliduser",
		Password: "invalidpassword",
	}
	reqPayload, _ := json.Marshal(loginReq)

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(reqPayload))
	assert.NoError(t, err)

	// Create a new Svc with the mock AuthsStorage
	authSvc := NewAuthSvc(storage, storage)

	// Call the LoginUser function
	_, statusCode, err := authSvc.LoginUser(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	// Create a mock AuthsStorage
	storage := mock.NewMockConnection()

	// Create Auth service
	authSvc := NewAuthSvc(storage, storage)

	// Create a sample user
	userDTO := &models.UserDTO{
		Username: "testuser",
		Pwd:      "password123",
		Email:    "sample@acme.com",
	}

	_ = authSvc.userStorage.AddUser(userDTO)
	sampleUser, _ := authSvc.userStorage.UserByName(userDTO.Username)

	// Create a JSON payload for the request with an invalid password
	loginReq := loginRequest{
		Username: sampleUser.Username,
		Password: "invalidpassword",
	}
	reqPayload, _ := json.Marshal(loginReq)

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(reqPayload))
	assert.NoError(t, err)

	// Call the LoginUser function
	_, statusCode, err := authSvc.LoginUser(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func TestParseCredentials(t *testing.T) {
	// Test with valid JSON
	username := "testuser"
	password := "testpassword"

	validJSON := `{"username": "testuser", "password": "testpassword"}`
	reader := io.NopCloser(strings.NewReader(validJSON))

	result, err := parseCredentials(reader)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, username, result.Username)
	assert.Equal(t, password, result.Password)

	// Test with Invalid JSON
	invalidJSON := `{"username": "testuser", "password": "testpassword"`
	reader = io.NopCloser(strings.NewReader(invalidJSON))

	result, err = parseCredentials(reader)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Test with empty JSON
	emptyJSON := `{}`
	reader = io.NopCloser(strings.NewReader(emptyJSON))

	result, err = parseCredentials(reader)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Test with valid JSON with extra field
	validJSONWithExtra := `{"username": "testuser", "password": "testpassword", "extra": "field"}`
	reader = io.NopCloser(strings.NewReader(validJSONWithExtra))

	result, err = parseCredentials(reader)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, username, result.Username)
	assert.Equal(t, password, result.Password)
}

func TestLoginUser_InvalidBody(t *testing.T) {
	// Create a JSON payload for the request
	loginReq := loginRequest{}
	reqPayload, _ := json.Marshal(loginReq)

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(reqPayload))
	assert.NoError(t, err)

	// Create a new Svc with the mock AuthsStorage
	authSvc := NewAuthSvc(nil, nil)

	// Call the LoginUser function
	_, statusCode, err := authSvc.LoginUser(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}

func TestLoginUser_FailToDeleteOldSession(t *testing.T) {
	// Create a mock AuthsStorage
	storage := mock.NewMockConnection()

	// Create Auth service
	authSvc := NewAuthSvc(storage, storage)

	// Create a sample user
	userDTO := &models.UserDTO{
		Username: "testuser",
		Pwd:      "password123",
		Email:    "sample@acme.com",
	}

	_ = authSvc.userStorage.AddUser(userDTO)
	sampleUser, _ := authSvc.userStorage.UserByName(userDTO.Username)

	// Create a JSON payload for the request
	loginReq := loginRequest{
		Username: sampleUser.Username,
		Password: sampleUser.Pwd,
	}
	reqPayload, _ := json.Marshal(loginReq)

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(reqPayload))
	assert.NoError(t, err)

	// Creating session in DB
	_, _, err = authSvc.LoginUser(req)
	assert.NoError(t, err)

	// Simulation of connection troubles
	storage.FailAuthMethod("DeleteAuth")
	defer storage.FixAuthMethod("DeleteAuth")

	// Creating of a new request
	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(reqPayload))

	// Call the LoginUser function
	_, statusCode, err := authSvc.LoginUser(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
}
