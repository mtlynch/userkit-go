# Go UserKit

## Summary
An experimental UserKit client library for Go.
NOTE: May be missing some features or be updated with breaking changes.

## Installation

```
go get github.com/workpail/workpail-go
```

## Documentation

For full examples and docs checkout [UserKit documentation][userkit-docs].

## Example usage

```go
package main

import (
	"fmt"
	userkit "github.com/workpail/userkit-go"
)

func main() {
	uk := userkit.NewUserKit("<YOUR_APP_SECRET_KEY>")

	// create a user
	data := map[string]string{"email": "jane.smith@example.com",
		"password": "secretpass"}
	user, err := uk.Users.Create(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", user)

	// login a user
	token, err := uk.Users.LoginUser("jane.smith@example.com", "secretpass", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", token)

	// fetch a logged-in user by their session-token
	user, err = uk.Users.GetCurrentUser(token.Token)
	if err != nil {
		fmt.Println("User not logged in. Error: ")
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user)
}
```

[userkit-docs]: https://docs.userkit.io
