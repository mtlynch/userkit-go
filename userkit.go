package userkit

import (
	"encoding/json"
	"io"
)

type UserKit struct {
	ukRq UKRequest
}

func NewUserKit(apiKey string) UserKit {
	ukr := UKRequest{ApiKey: apiKey}
	userkit := UserKit{ukRq: ukr}
	return userkit
}

func (uk *UserKit) GetCurrentUser(sessionToken string) (*User, error) {
	r, err := uk.ukRq.Do("GET", "https://api.userkit.io/v1/users/by_token", nil, &sessionToken)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, processErrResp(r.Body)
	}

	var user User
	json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uk *UserKit) LoginUser(username, password string) (*SessionToken, error) {
	type payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	data := payload{Username: username, Password: password}
	r, err := uk.ukRq.Do("POST", "https://api.userkit.io/v1/users/login", data, nil)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, processErrResp(r.Body)
	}

	var token SessionToken
	json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

type User struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	VerifiedEmail   string  `json:"verified_email"`
	VerifiedPhone   string  `json:"verified_phone"`
	AuthType        string  `json:"auth_type"`
	LastFailedLogin float64 `json:"last_failed_login"`
	LastLogin       float64 `json:"last_login"`
	Disabled        string  `json:"disabled"`
	Created         float64 `json:"created"`
}

type SessionToken struct {
	Token            string  `json:"token"`
	ExpiresInSecs    float64 `json:"expires_in_secs"`
	RefreshAfterSecs float64 `json:"refresh_after_secs"`
}

type UserKitError struct {
	ErrorValue struct {
		Type      string  `json:"type"`
		Code      string  `json:"code"`
		Message   string  `json:"message"`
		RetryWait float64 `json:"retry_wait"`
	} `json:"error"`
}

func (e UserKitError) Error() string { return e.ErrorValue.Message }

// processErrResp takes a JSON UserKit error string and returns
// a UserKitError object.
func processErrResp(body io.ReadCloser) error {
	var ukError UserKitError
	json.NewDecoder(body).Decode(&ukError)
	return ukError
}
