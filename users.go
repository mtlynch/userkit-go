package userkit

import (
	"encoding/json"
)

type usersClient struct {
	c client
}

func (c *usersClient) Create(data map[string]string) (*User, error) {
	rq := c.c.ukRq
	r, err := rq.Do("POST", apiURL+"/users", data, nil)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, processErrListResp(r.Body)
	}

	var user User
	err = json.Unmarshal(r.Body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *usersClient) Get(userId string) (*User, error) {
	rq := c.c.ukRq
	url := apiURL + "/users/" + userId
	r, err := rq.Do("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, processErrResp(r.Body)
	}

	var user User
	err = json.Unmarshal(r.Body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *usersClient) GetCurrentUser(sessionToken string) (*User, error) {
	rq := c.c.ukRq
	r, err := rq.Do("GET", apiURL+"/users/by_token", nil, &sessionToken)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, processErrResp(r.Body)
	}

	var user User
	err = json.Unmarshal(r.Body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *usersClient) Login(username, password, loginCode string) (*SessionToken, error) {
	rq := c.c.ukRq
	type payload struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		LoginCode string `json:"login_code"`
	}
	data := payload{Username: username}
	if password != "" {
		data.Password = password
	}
	if loginCode != "" {
		data.LoginCode = loginCode
	}
	r, err := rq.Do("POST", apiURL+"/users/login", data, nil)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, processErrResp(r.Body)
	}

	var token SessionToken
	err = json.Unmarshal(r.Body, &token)
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
	Disabled        bool    `json:"disabled"`
	Created         float64 `json:"created"`
}

type SessionToken struct {
	Token            string  `json:"token"`
	ExpiresInSecs    float64 `json:"expires_in_secs"`
	RefreshAfterSecs float64 `json:"refresh_after_secs"`
}
