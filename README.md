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
	data := map[string]string{
		"email":    "jane.smith@example.com",
		"password": "secretpass"}

	user, err := uk.Users.Create(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("USER: %+v", user)

	// fetch a user
	user, _ = uk.Users.Get("usr_j3LB5QPAH8B9UD")

	// update a user
	upData := map[string]string{"name": "Jane Smith"}
	user, _ = uk.Users.Update("usr_j3LB5QPAH8B9UD", upData)

	// login a user
	token, _ := uk.Users.Login("jane.smith@example.com", "secretpass", "")

	// fetch a logged-in user by their session-token
	user, err = uk.Users.GetCurrentUser(token.Token)
	if err != nil {
		fmt.Println("User not logged in. Error: ")
		fmt.Println(err)
		return
	}
	fmt.Printf("GET USER BY SESSION: %+v\n", user)
}
```

## Running tests

To run tests you need to create a test-app.

Set the `USERKIT_KEY` environment variable to your test app key, then run go test:
```
USERKIT_KEY=<YOUR_APP_SECRET_KEY> go test -v
```

[userkit-docs]: https://docs.userkit.io
