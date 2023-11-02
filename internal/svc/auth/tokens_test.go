package auth

import (
	"github.com/stretchr/testify/assert"
	"mzda/internal/storage/db/mock"
	"mzda/internal/storage/models"
	"testing"
)

func TestGenerateTokens_Success(t *testing.T) {
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

	// Call the generateTokens function
	jwt, refresh, err := authSvc.generateTokens(sampleUser)
	assert.NoError(t, err)

	// Check if the JWT and Refresh token were generated
	assert.NotEqual(t, "", jwt)
	assert.NotEqual(t, "", refresh)

	storedSession, err := storage.GetAuth(refresh)
	assert.NoError(t, err)
	assert.Equal(t, storedSession.RefreshToken, refresh)
	assert.Equal(t, storedSession.Username, sampleUser.Username)
}

func TestGenerateTokens_FailedStoreSession(t *testing.T) {
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

	// Call the generateTokens function
	storage.FailAuthMethod("AddAuth")
	defer storage.FixAuthMethod("AddAuth")
	_, _, err := authSvc.generateTokens(sampleUser)
	assert.Error(t, err)
}
