package userkit

// Client is a UserKit client through which sub-resources can be
// accessed
type Client struct {
	ukRq Requestor
	key  string

	// resources
	Users *usersClient
}

const (
	apiURL = "https://api.userkit.io/v1"
)

// NewUserKit creates a new client
func NewUserKit(apiKey string) Client {
	r := Requestor{APIKey: apiKey}

	uk := Client{ukRq: r, key: apiKey}
	uk.Users = &usersClient{c: uk}
	return uk
}
