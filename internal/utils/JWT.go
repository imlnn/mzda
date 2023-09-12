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
	const fn = "internal/auth/utils/JWT/newHeader"
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
	Admin    bool   `json:"admin"`
}

func newPayload(username string, admin bool) *payload {
	const fn = "internal/auth/utils/JWT/newPayload"
	//iss := os.Getenv("SVC")
	iss := "MZDA_AUTH_SVC"
	return &payload{Iss: iss,
		Iat:      time.Now().Unix(),
		Exp:      time.Now().Add(30 * time.Minute).Unix(),
		Username: username,
		Admin:    admin}
}

type JWT struct {
	token    string
	exp      time.Time
	username string
	admin    bool
}

//func NewJWT(jwt string) (*JWT, error) {
//	if !ValidateJWT(jwt) {
//		return nil, fmt.Errorf("jwt signature validation failed")
//	}
//	payload, err := decodeJWTPayload(jwt)
//	if err != nil {
//		return nil, err
//	}
//	token := JWT{token: jwt,
//		exp:      time.Unix(payload.Exp, 0),
//		username: payload.Username,
//		admin:    payload.Admin}
//
//	if !token.IsExpired() {
//		return nil, fmt.Errorf("jwt is expired")
//	}
//
//	return &token, nil
//}

func (t *JWT) Username() string {
	const fn = "internal/auth/utils/JWT/Username"
	return t.username
}

func (t *JWT) Admin() bool {
	const fn = "internal/auth/utils/JWT/Admin"
	return t.admin
}

func (t *JWT) IsExpired() bool {
	const fn = "internal/auth/utils/JWT/IsExpired"
	return t.exp.After(time.Now())
}

func GenerateJWT(username string, role models.Role) (string, error) {
	const fn = "internal/auth/utils/JWT/GenerateJWT"
	//secret := os.Getenv("jwtSecret")
	secret := "secret"

	var admin = false
	if role == models.ADMIN {
		admin = true
	}

	header := newHeader(defaultAlg, defaultTyp)
	payload := newPayload(username, admin)

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

func ValidateJWT(token string) bool {
	const fn = "internal/auth/utils/JWT/ValidateJWT"

	//secret := os.Getenv("jwtSecret")
	secret := "secret"
	data := strings.Split(token, ".")
	hp := data[0] + "." + data[1]
	enc := hmac.New(sha512.New, []byte(secret))
	enc.Write([]byte(hp))
	signature := hex.EncodeToString(enc.Sum(nil))
	return strings.EqualFold(data[2], signature)
}

func decodeJWTPayload(token string) (*payload, error) {
	const fn = "internal/auth/utils/JWT/decodeJWTPayload"

	data := strings.Split(token, ".")
	p, err := base64.RawStdEncoding.DecodeString(data[1])
	fmt.Println(string(p))
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}
	var payload payload
	err = json.Unmarshal(p, &payload)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}
	return &payload, nil
}
