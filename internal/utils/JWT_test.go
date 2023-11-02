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
	username := "testuser"
	userID := 1
	role := models.USER

	token, err := GenerateJWT(username, userID, role)
	assert.NoError(t, err)

	jwt, err := NewJWT(token)
	assert.NoError(t, err)

	assert.Equal(t, username, jwt.Username)
	assert.Equal(t, userID, jwt.UserID)
	assert.False(t, jwt.Admin)
}

func TestIsInvalidJWT(t *testing.T) {
	username := "testuser"
	userID := 1
	role := models.USER

	invalidToken := "invalid.token"
	validToken, _ := GenerateJWT(username, userID, role)
	invalidSignatureToken := validToken[:len(validToken)-5]

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

	var payload payload

	err = json.Unmarshal(decodedPayload, &payload)
	payload.Exp = time.Now().Unix() - 3600
	newPayloadJSON, _ := json.Marshal(payload)
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
