package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type payload struct {
	Iss      string `json:"iss"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
}

type Token struct {
	token    string
	exp      time.Time
	username string
	admin    bool
}

func NewToken(jwt string) (*Token, error) {
	if !ValidateJWT(jwt) {
		return nil, fmt.Errorf("jwt signature validation failed")
	}
	payload, err := decodeJWTPayload(jwt)
	if err != nil {
		return nil, err
	}
	token := Token{token: jwt,
		exp:      time.Unix(payload.Exp, 0),
		username: payload.Username,
		admin:    payload.Admin}

	if !token.IsExpired() {
		return nil, fmt.Errorf("jwt is expired")
	}

	return &token, nil
}

func (t *Token) Username() string {
	return t.username
}

func (t *Token) Admin() bool {
	return t.admin
}

func (t *Token) IsExpired() bool {
	return t.exp.After(time.Now())
}

func GenerateJWT(username string, issuer string, admin bool) ([]byte, error) {
	//secret := os.Getenv("jwtSecret")
	secret := "secret"

	header := header{Alg: "HS512",
		Typ: "JWT"}

	payload := payload{Iss: issuer,
		Iat:      time.Now().Unix(),
		Exp:      time.Now().Add(30 * time.Minute).Unix(),
		Username: username,
		Admin:    admin}

	var headerJSON []byte
	var payloadJSON []byte

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return []byte(""), err
	}

	payloadJSON, err = json.Marshal(payload)
	if err != nil {
		return []byte(""), err
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
	return token, nil
}

func ValidateJWT(token string) bool {
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
	data := strings.Split(token, ".")
	p, err := base64.RawStdEncoding.DecodeString(data[1])
	fmt.Println(string(p))
	if err != nil {
		return nil, err
	}
	var payload payload
	err = json.Unmarshal(p, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
