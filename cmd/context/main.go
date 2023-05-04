package main

import (
	stdctx "context"
	"fmt"

	"github.com/gsom95/go-web-dev/context"
	"github.com/gsom95/go-web-dev/models"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "jon@calhoun.io",
	}
	ctx = context.WithUser(ctx, &user)

	retrievedUser := context.User(ctx)
	fmt.Println(retrievedUser.Email)
}
