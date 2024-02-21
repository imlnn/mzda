package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mzda/internal/storage/db/postgres"
	"mzda/internal/storage/models"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const HOST = "http://127.0.0.1:32000/api/v1.0"

// Flushing DB after tests
func flushDB() {
	connection, err := postgres.New()
	if err != nil {
		return
	}
	connection.FlushDB()
}

func BuildUserDTO(reqNum int) *models.UserDTO {
	return &models.UserDTO{
		Username: "LoadTestUser" + strconv.Itoa(reqNum),
		Pwd:      "TestUserPassword" + strconv.Itoa(reqNum),
		Email:    "test" + strconv.Itoa(reqNum) + "@acme.com",
	}
}

func BuildCreateUserRequest(dto *models.UserDTO) *http.Request {
	reqPayload, _ := json.Marshal(dto)

	req, _ := http.NewRequest("POST", HOST+"/signup", bytes.NewReader(reqPayload))
	return req
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWT     string `json:"jwt"`
	Refresh string `json:"refresh"`
}

func BuildLoginRequest(user *models.UserDTO) *http.Request {
	loginReq := LoginRequest{
		Username: user.Username,
		Password: user.Pwd,
	}
	reqPayload, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", HOST+"/auth/signin", bytes.NewReader(reqPayload))
	return req
}

func DecodeLoginResponse(res *http.Response) *LoginResponse {
	var tokens LoginResponse
	_ = json.NewDecoder(res.Body).Decode(&tokens)
	return &tokens
}

func BuildRenewRequest(refresh string) *http.Request {
	req, _ := http.NewRequest("POST", HOST+"/auth/renew", nil)
	req.Header.Add("refreshToken", refresh)
	return req
}

type ChangeEmailRequest struct {
	Username string `json:"username"`
	NewEmail string `json:"newEmail"`
}

func BuildChangeEmailRequest(dto *models.UserDTO, tokens *LoginResponse) *http.Request {
	emailReq := ChangeEmailRequest{
		Username: dto.Username,
		NewEmail: "changed" + dto.Email,
	}
	reqPayload, _ := json.Marshal(emailReq)

	req, _ := http.NewRequest("POST", HOST+"/user/changeEmail", bytes.NewReader(reqPayload))
	req.Header.Add("Authorization", tokens.JWT)
	return req
}

type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func BuildChangePasswordRequest(dto *models.UserDTO, tokens *LoginResponse) *http.Request {
	pwdReq := ChangePasswordRequest{
		Username:    dto.Username,
		OldPassword: dto.Pwd,
		NewPassword: "changed" + dto.Pwd,
	}
	reqPayload, _ := json.Marshal(pwdReq)

	req, _ := http.NewRequest("POST", HOST+"/user/changePassword", bytes.NewReader(reqPayload))
	req.Header.Add("Authorization", tokens.JWT)
	return req
}

type ChangeUsernameRequest struct {
	Username    string `json:"username"`
	NewUsername string `json:"newUsername"`
}

func BuildChangeUsernameRequest(dto *models.UserDTO, tokens *LoginResponse) *http.Request {
	usernameReq := ChangeUsernameRequest{
		Username:    dto.Username,
		NewUsername: "changed" + dto.Username,
	}
	reqPayload, _ := json.Marshal(usernameReq)

	req, _ := http.NewRequest("POST", HOST+"/user/changeUsername", bytes.NewReader(reqPayload))
	req.Header.Add("Authorization", tokens.JWT)
	return req
}

func main() {
	flushDB()
	totalRequests := 1000
	rps := 100

	for {
		// Channels for transmitting errors from gorutines
		createUserErrorsChan := make(chan error, totalRequests)
		loginErrorsChan := make(chan error, totalRequests)
		renewErrorsChan := make(chan error, totalRequests)
		changeEmailErrorsChan := make(chan error, totalRequests)
		changePasswordErrorsChan := make(chan error, totalRequests)
		changeUsernameErrorsChan := make(chan error, totalRequests)

		// Stores errors count
		createUserErrors := 0
		loginErrors := 0
		renewErrors := 0
		changeEmailErrors := 0
		changePasswordErrors := 0
		changeUsernameErrors := 0

		var wg sync.WaitGroup
		wg.Add(totalRequests)
		for i := 0; i < totalRequests; i++ {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Println("Panic:", r)
						wg.Done()
						return
					}
				}()

				// Registration
				dto := BuildUserDTO(i)
				regReq := BuildCreateUserRequest(dto)
				client := http.Client{
					Timeout: time.Second * 5,
				}
				res, err := client.Do(regReq)
				if err != nil {
					log.Printf("CreateUser: %v, Response code: %v, Failed with err: %v", i, res.StatusCode, err)
				}
				log.Printf("CreateUser: %v, Response code: %v, Status: %v", i, res.StatusCode, res.Status)
				if res.StatusCode != http.StatusOK {
					createUserErrorsChan <- err
					wg.Done()
					return
				}

				err = res.Body.Close()
				if err != nil {
					return
				}

				// Login
				loginReq := BuildLoginRequest(dto)
				res, err = client.Do(loginReq)
				if err != nil {
					log.Printf("Login: %v, Response code: %v, Failed with err: %v", i, res.StatusCode, err)
				}
				log.Printf("Login: %v, Response code: %v, Status: %v", i, res.StatusCode, res.Status)
				if res.StatusCode != http.StatusOK {
					loginErrorsChan <- err
					wg.Done()
					return
				}
				tokens := DecodeLoginResponse(res)

				err = res.Body.Close()
				if err != nil {
					return
				}

				//Token renewal
				renewReq := BuildRenewRequest(tokens.Refresh)
				res, err = client.Do(renewReq)
				if err != nil {
					log.Printf("Renew: %v, Response code: %v, Failed with err: %v", i, res.StatusCode, err)
				}
				log.Printf("Renew: %v, Response code: %v, Status: %v", i, res.StatusCode, res.Status)
				if res.StatusCode != http.StatusOK {
					renewErrorsChan <- err
					wg.Done()
					return
				}
				tokens = DecodeLoginResponse(res)

				err = res.Body.Close()
				if err != nil {
					return
				}

				// Changing of email
				emailReq := BuildChangeEmailRequest(dto, tokens)
				res, err = client.Do(emailReq)
				if err != nil {
					log.Printf("ChangeEmail: %v, Response code: %v, Failed with err: %v", i, res.StatusCode, err)
				}
				log.Printf("ChangeEmail: %v, Response code: %v, Status: %v", i, res.StatusCode, res.Status)
				if res.StatusCode != http.StatusOK {
					changeEmailErrorsChan <- err
					wg.Done()
					return
				}
				dto.Email = "changed" + dto.Email

				err = res.Body.Close()
				if err != nil {
					return
				}

				// Changing of password
				pwdReq := BuildChangePasswordRequest(dto, tokens)
				res, err = client.Do(pwdReq)
				if err != nil {
					log.Printf("ChangePassword: %v, Response code: %v, Failed with err: %v", i, res.StatusCode, err)
				}
				log.Printf("ChangePassword: %v, Response code: %v, Status: %v", i, res.StatusCode, res.Status)
				if res.StatusCode != http.StatusOK {
					changePasswordErrorsChan <- err
					wg.Done()
					return
				}
				dto.Pwd = "changed" + dto.Pwd

				err = res.Body.Close()
				if err != nil {
					return
				}

				// Changing of username
				usernameReq := BuildChangeUsernameRequest(dto, tokens)
				res, err = client.Do(usernameReq)
				if err != nil {
					log.Printf("ChangeUsername: %v, Response code: %v, Failed with err: %v", i, res.StatusCode, err)
				}
				log.Printf("ChangeUsername: %v, Response code: %v, Status: %v", i, res.StatusCode, res.Status)
				if res.StatusCode != http.StatusOK {
					changeUsernameErrorsChan <- err
					wg.Done()
					return
				}
				dto.Username = "changed" + dto.Username

				err = res.Body.Close()
				if err != nil {
					return
				}

				wg.Done()
				return
			}()
			time.Sleep(time.Second / time.Duration(rps))
		}

		wg.Wait()
		// Count errors by endpoints
		for i := 0; i < totalRequests; i++ {
			select {
			case <-createUserErrorsChan:
				createUserErrors++
			default:
			}

			select {
			case <-loginErrorsChan:
				loginErrors++
			default:
			}

			select {
			case <-renewErrorsChan:
				renewErrors++
			default:
			}

			select {
			case <-changeEmailErrorsChan:
				changeEmailErrors++
			default:
			}

			select {
			case <-changePasswordErrorsChan:
				changePasswordErrors++
			default:
			}

			select {
			case <-changeUsernameErrorsChan:
				changeUsernameErrors++
			default:
			}
		}

		close(createUserErrorsChan)
		close(loginErrorsChan)
		close(renewErrorsChan)
		close(changeEmailErrorsChan)
		close(changePasswordErrorsChan)
		close(changeUsernameErrorsChan)

		totalErrors := createUserErrors + loginErrors + renewErrors + changeEmailErrors + changePasswordErrors + changeUsernameErrors
		successRate := (float32(totalRequests) - float32(totalErrors)) / float32(totalRequests)

		// Getting stats
		fmt.Printf("\nLoad test report\nTotal requests: %v\nRPS: %v\nErrors: %v\nSuccessRate: %v",
			totalRequests, rps, totalErrors, successRate)
		fmt.Printf("\n\nErrors by endpoint\nCreateUser: %v\nSignIn: %v\nRenew: %v\nChangeEmail: %v\nChangePassword: %v\nChangeUsername:%v\n",
			createUserErrors, loginErrors, renewErrors, changeEmailErrors, changePasswordErrors, changeUsernameErrors)

		time.Sleep(5 * time.Second)

		// If more than 90% of the requests were successful test continues with additional 10 RPS
		// Else test stopping
		if successRate > 0.9 {
			totalRequests = totalRequests + 100
			rps = totalRequests / 10
			flushDB()
		} else {
			flushDB()
			break
		}
	}
}
