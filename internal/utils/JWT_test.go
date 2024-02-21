package utils

import (
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"mzda/internal/storage/models"
	"strings"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	// Creating sample payload data
	username := "testuser"
	userID := 1
	role := models.USER

	// Generating JWT token
	token, err := GenerateJWT(username, userID, role)
	assert.NoError(t, err)

	// Parsing token to JWT struct
	jwt, err := NewJWT(token)
	assert.NoError(t, err)

	// Check equality of test data and test data after encoding
	assert.Equal(t, username, jwt.Username)
	assert.Equal(t, userID, jwt.UserID)
	assert.False(t, jwt.Admin)
}

func TestIsInvalidJWT(t *testing.T) {
	// Creating sample payload data
	username := "testuser"
	userID := 1
	role := models.USER

	// Creating not JWT string, valid JWT, and JWT with invalid signature
	invalidToken := "invalid.token"
	validToken, _ := GenerateJWT(username, userID, role)
	invalidSignatureToken := validToken[:len(validToken)-5]

	// Checking IsInvalidJWT
	assert.True(t, IsInvalidJWT(invalidToken))
	assert.True(t, IsInvalidJWT(invalidSignatureToken))
	assert.False(t, IsInvalidJWT(validToken))
}

func TestDecodePayload(t *testing.T) {
	username := "testuser"
	userID := 1
	role := models.USER

	token, _ := GenerateJWT(username, userID, role)

	_, err := NewJWT(token)
	assert.NoError(t, err)

	expiredToken, _ := GenerateJWT(username, userID, role)

	jwtParts := strings.Split(expiredToken, ".")
	decodedPayload, _ := base64.RawStdEncoding.DecodeString(jwtParts[1])

	var pl payload

	_ = json.Unmarshal(decodedPayload, &pl)
	pl.Exp = time.Now().Unix() - 3600
	newPayloadJSON, _ := json.Marshal(pl)
	newPayload := base64.RawStdEncoding.EncodeToString(newPayloadJSON)
	expiredToken = jwtParts[0] + "." + newPayload + "." + jwtParts[2]

	_, err = NewJWT(expiredToken)
	assert.Error(t, err)
}

func TestIsExpired(t *testing.T) {
	normalJWT := JWT{
		Token:    "",
		Exp:      time.Now().Add(30 * time.Minute),
		Username: "",
		UserID:   0,
		Admin:    false,
	}

	invalidJWT := JWT{
		Token:    "",
		Exp:      time.Now().Add(-30 * time.Minute),
		Username: "",
		UserID:   0,
		Admin:    false,
	}

	assert.False(t, normalJWT.IsExpired())
	assert.True(t, invalidJWT.IsExpired())
}

func TestDecodeJWTPayload(t *testing.T) {
	username := "testuser"
	userID := 1
	role := models.USER

	token, _ := GenerateJWT(username, userID, role)

	payload, err := decodeJWTPayload(token)
	assert.NoError(t, err)

	assert.Equal(t, username, payload.Username)
	assert.Equal(t, userID, payload.UserID)
	assert.False(t, payload.Admin)
}
