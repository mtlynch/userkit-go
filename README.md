# userkit-go
A UserKit client library for Go

# Installation

```
go get github.com/workpail/workpail-go
```

# Documentation

For full examples and docs checkout [UserKit documentation][userkit-docs].

# Example usage

```
package main

import (
	"fmt"
	userkit "github.com/workpail/userkit-go"
)

func main() {
	uk := userkit.NewUserKit("<YOUR_APP_SECRET_KEY>")

	token, err := uk.LoginUser("jane.smith@example.com", "password")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", token)

	user, err := uk.GetCurrentUser(token.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user)
}

```

[userkit-docs]: https://docs.userkit.io
