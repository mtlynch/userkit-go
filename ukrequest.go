package userkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// UKRequest holds info needed to make an authenticated request
// to the UserKit API endpoints.
type UKRequest struct {
	ApiKey string
}

// Do makes an authenticated request to the UserKit api. Make
// sure you replace <USERKIT_APP_KEY> with your app key. payload will
// be sent as a json request body. Some requests (such as to the
// users/by_token endpoint) require a sessionToken parameter.
func (ukrequest *UKRequest) Do(method, url string, payload interface{},
	sessionToken *string) (*http.Response, error) {
	client := &http.Client{}
	b := new(bytes.Buffer)

	if payload != nil {
		json.NewEncoder(b).Encode(payload)
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if sessionToken != nil {
		req.Header.Set("X-User-Token", *sessionToken)
	}
	req.SetBasicAuth("api", ukrequest.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("\nERROR: %v", resp)
		fmt.Println(err.Error())
		return nil, err
	}
	return resp, nil
}
