package userkit

import (
	"encoding/json"
	"fmt"
	"strings"
)

type UserKitError struct {
	Type      string  `json:"type"`
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	RetryWait float64 `json:"retry_wait"`
	Param     string  `json:"param"`
}

func (e UserKitError) Error() string { return e.Message }

type UserKitErrorList struct {
	Errors []UserKitError
}

func (e UserKitErrorList) Error() string {
	msgs := make([]string, 0)
	for _, err := range e.Errors {
		msg := fmt.Sprintf("\"%s\"", err.Message)
		msgs = append(msgs, msg)
	}
	return fmt.Sprintf("userkit errors (%d): [%s]",
		len(e.Errors), strings.Join(msgs, ", "))
}

type userkitErrorResponse struct {
	ErrorValue UserKitError `json:"error"`
}

type userkitErrorListResponse struct {
	Errors []userkitErrorResponse `json:"errors"`
}

// processErrResp takes a JSON UserKit error string and returns
// a UserKitError object.
func processErrResp(body []byte) error {
	var ukErrResp userkitErrorResponse
	json.Unmarshal(body, &ukErrResp)
	var ukError UserKitError
	ukError = ukErrResp.ErrorValue
	return ukError
}

func processErrListResp(body []byte) error {
	var ukErrListResp userkitErrorListResponse
	json.Unmarshal(body, &ukErrListResp)
	errs := make([]UserKitError, 0)
	for _, e := range ukErrListResp.Errors {
		errs = append(errs, e.ErrorValue)
	}
	return UserKitErrorList{Errors: errs}
}
