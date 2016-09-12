package userkit

type UserKit struct {
	Users *usersClient
}

type client struct {
	ukRq UKRequest
	key  string
}

const (
	apiURL = "https://api.userkit.io/v1"
)

func NewUserKit(apiKey string) UserKit {
	r := UKRequest{ApiKey: apiKey}
	c := client{ukRq: r, key: apiKey}

	uk := UserKit{}
	uk.Users = &usersClient{c: c}
	return uk
}
