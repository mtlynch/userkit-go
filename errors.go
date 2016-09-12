package userkit

type UserKitError struct {
	Type      string  `json:"type"`
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	RetryWait float64 `json:"retry_wait"`
	Param     string  `json:"param"`
}

func (e UserKitError) Error() string { return e.Message }
