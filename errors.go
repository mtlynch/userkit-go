package userkit

import "encoding/json"

// CoreError is a base UserKit exception
type CoreError struct {
	Type      string  `json:"type"`
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	RetryWait float64 `json:"retry_wait"`
	Param     string  `json:"param"`
}

// Error is what developers should usually expect to handle
// when using this SDK. It contains an Errors property which is
// populated with a list of CoreError's when hitting an endpoint
// which can return multiple errors.
type Error struct {
	CoreError
	Errors []CoreError
}

func (e CoreError) Error() string { return e.Message }

type userkitErrorResponse struct {
	ErrorValue CoreError   `json:"error"`
	Errors     []CoreError `json:"errors"`
}

// processErrResp takes a JSON UserKit error string and returns
// a UserKitError object.
func processErrResp(body []byte) Error {
	var errResp userkitErrorResponse
	json.Unmarshal(body, &errResp)
	e := errResp.ErrorValue

	err := Error{}
	err.Type = e.Type
	err.Code = e.Code
	err.Message = e.Message
	err.RetryWait = e.RetryWait
	err.Param = e.Param
	err.Errors = errResp.Errors

	return err
}
