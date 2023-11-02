package auth

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"mzda/internal/storage/db/mock"
	"mzda/internal/storage/models"
	"net/http"
	"testing"
	"time"
)

func TestRenew_Success(t *testing.T) {
	// Create a mock AuthsStorage
	storage := &mock.Connection{}

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

	// Generate a sample JWT and Refresh token
	_, refresh, _ := authSvc.generateTokens(sampleUser)

	// Create a request with the JSON payload
	req, _ := http.NewRequest("POST", "/renew", bytes.NewReader(nil))

	req.Header.Add("refreshToken", refresh)

	// Call the Renew function
	response, statusCode, err := authSvc.Renew(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	// Decode the response
	var renewRes renewResponse
	err = json.Unmarshal(response, &renewRes)
	assert.NoError(t, err)

	// Check if the JWT and Refresh token were generated
	assert.NotEqual(t, "", renewRes.JWT)
	assert.NotEqual(t, "", renewRes.Refresh)

	storedSession, err := storage.GetAuth(renewRes.Refresh)
	assert.NoError(t, err)
	assert.Equal(t, storedSession.RefreshToken, renewRes.Refresh)
	assert.Equal(t, storedSession.Username, sampleUser.Username)
}

func TestRenew_MissingRefreshToken(t *testing.T) {
	// Create a mock AuthsStorage
	storage := &mock.Connection{}

	// Create a new Svc with the mock AuthsStorage
	authSvc := NewAuthSvc(storage, storage)

	// Create a request with the JSON payload
	req, _ := http.NewRequest("POST", "/renew", bytes.NewReader(nil))

	// Call the Renew function
	_, statusCode, err := authSvc.Renew(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)
}

func TestRenew_ExpiredRefreshToken(t *testing.T) {
	// Create a mock AuthsStorage
	storage := &mock.Connection{}

	// Create Auth service
	authSvc := NewAuthSvc(storage, storage)

	userDTO := &models.UserDTO{
		Username: "testuser",
		Pwd:      "password123",
		Email:    "sample@acme.com",
	}

	_ = authSvc.userStorage.AddUser(userDTO)
	sampleUser, _ := authSvc.userStorage.UserByName(userDTO.Username)

	// Generate a sample Refresh token
	_, refresh, _ := authSvc.generateTokens(sampleUser)

	// Create a request
	req, _ := http.NewRequest("POST", "/renew", bytes.NewReader(nil))
	req.Header.Add("refreshToken", refresh)

	// Add the generated Auth object to the mock AuthsStorage
	_ = authSvc.authStorage.DeleteAuth(refresh)
	_ = authSvc.authStorage.AddAuth(&models.Auth{
		Username:     sampleUser.Username,
		RefreshToken: refresh,
		Expires:      time.Now().AddDate(0, -1, 0),
	})

	// Call the Renew function
	_, statusCode, err := authSvc.Renew(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func TestRenew_AuthNotFound(t *testing.T) {
	// Create a mock AuthsStorage
	storage := &mock.Connection{}

	// Create Auth service
	authSvc := NewAuthSvc(storage, storage)

	// Generate a sample JWT and Refresh token
	refresh := "notexistentoken"

	// Create a request
	req, _ := http.NewRequest("POST", "/renew", bytes.NewReader(nil))
	req.Header.Add("refreshToken", refresh)

	// Call the Renew function
	_, statusCode, err := authSvc.Renew(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestRenew_UserNotExists(t *testing.T) {
	// Create a mock AuthsStorage
	storage := &mock.Connection{}

	// Create Auth service
	authSvc := NewAuthSvc(storage, storage)

	userDTO := &models.UserDTO{
		Username: "testuser",
		Pwd:      "password123",
		Email:    "sample@acme.com",
	}

	_ = authSvc.userStorage.AddUser(userDTO)
	sampleUser, _ := authSvc.userStorage.UserByName(userDTO.Username)

	// Generate a sample Refresh token
	_, refresh, _ := authSvc.generateTokens(sampleUser)

	// Create a request
	req, _ := http.NewRequest("POST", "/renew", bytes.NewReader(nil))
	req.Header.Add("refreshToken", refresh)

	// Add the generated Auth object to the mock AuthsStorage
	_ = authSvc.authStorage.AddAuth(&models.Auth{
		Username:     sampleUser.Username,
		RefreshToken: refresh,
		Expires:      time.Now().Add(time.Hour),
	})

	_ = authSvc.userStorage.DeleteUser(sampleUser)

	// Call the Renew function
	_, statusCode, err := authSvc.Renew(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestRenew_FailedDeleteOfOldSession(t *testing.T) {
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

	// Generate a sample JWT and Refresh token
	_, refresh, _ := authSvc.generateTokens(sampleUser)

	// Create a request with the JSON payload
	req, _ := http.NewRequest("POST", "/renew", bytes.NewReader(nil))
	req.Header.Add("refreshToken", refresh)

	// Call the Renew function
	storage.FailAuthMethod("DeleteAuth")
	defer storage.FixAuthMethod("DeleteAuth")
	_, statusCode, err := authSvc.Renew(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
}
