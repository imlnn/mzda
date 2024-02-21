package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"mzda/internal/storage/models"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultAlg = "HS512"
	defaultTyp = "JWT"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func newHeader(alg string, typ string) *header {
	const fn = "internal/utils/JWT/newHeader"

	return &header{
		Alg: alg,
		Typ: typ,
	}
}

type payload struct {
	Iss      string `json:"iss"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	Admin    bool   `json:"admin"`
}

func newPayload(username string, userID int, admin bool) *payload {
	const fn = "internal/utils/JWT/newPayload"
	const iss = "MZDA_AUTH_SVC"

	var tokenTTL, err = strconv.Atoi(os.Getenv("tokenTTL"))
	if err != nil {
		log.Printf("%v Couldn't get token TTL for enviroment", fn)
		return nil
	}

	return &payload{Iss: iss,
		Iat:      time.Now().Unix(),
		Exp:      time.Now().Add(time.Duration(tokenTTL) * time.Minute).Unix(),
		Username: username,
		UserID:   userID,
		Admin:    admin}
}

type JWT struct {
	Token    string
	Exp      time.Time
	Username string
	UserID   int
	Admin    bool
}

func NewJWT(jwt string) (*JWT, error) {
	const fn = "internal/utils/JWT/NewJWT"

	if IsInvalidJWT(jwt) {
		return nil, fmt.Errorf("jwt signature validation failed")
	}

	payload, err := decodeJWTPayload(jwt)
	if err != nil {
		return nil, err
	}
	token := JWT{Token: jwt,
		Exp:      time.Unix(payload.Exp, 0),
		Username: payload.Username,
		UserID:   payload.UserID,
		Admin:    payload.Admin}

	if token.IsExpired() {
		return nil, fmt.Errorf("jwt is expired")
	}

	return &token, nil
}

func (t *JWT) IsExpired() bool {
	const fn = "internal/utils/JWT/IsExpired"
	return t.Exp.Before(time.Now())
}

func GenerateJWT(username string, userID int, role models.Role) (string, error) {
	const fn = "internal/utils/JWT/GenerateJWT"
	secret := os.Getenv("jwtSecret")

	var admin = false
	if role == models.ADMIN {
		admin = true
	}

	header := newHeader(defaultAlg, defaultTyp)
	payload := newPayload(username, userID, admin)

	var headerJSON []byte
	var payloadJSON []byte

	headerJSON, err := json.Marshal(header)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return "", err
	}

	payloadJSON, err = json.Marshal(payload)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return "", err
	}

	var signature []byte
	var token []byte
	var h []byte
	var p []byte

	h = []byte(base64.RawStdEncoding.EncodeToString(headerJSON))
	p = []byte(base64.RawStdEncoding.EncodeToString(payloadJSON))

	token = append(token, h...)
	token = append(token, '.')
	token = append(token, p...)

	enc := hmac.New(sha512.New, []byte(secret))
	enc.Write(token)
	signature = []byte(hex.EncodeToString(enc.Sum(nil)))
	token = append(token, '.')
	token = append(token, signature...)

	return string(token), nil
}

func IsInvalidJWT(token string) bool {
	const fn = "internal/utils/JWT/IsInvalidJWT"
	const tokenParts = 3

	secret := os.Getenv("jwtSecret")
	data := strings.Split(token, ".")
	if len(data) != tokenParts {
		return true
	}

	hp := data[0] + "." + data[1]
	enc := hmac.New(sha512.New, []byte(secret))
	enc.Write([]byte(hp))
	signature := hex.EncodeToString(enc.Sum(nil))
	return !strings.EqualFold(data[2], signature)
}

func decodeJWTPayload(token string) (*payload, error) {
	const fn = "internal/utils/JWT/decodeJWTPayload"

	data := strings.Split(token, ".")
	p, err := base64.RawStdEncoding.DecodeString(data[1])
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}
	var pl payload
	err = json.Unmarshal(p, &pl)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}
	return &pl, nil
}
