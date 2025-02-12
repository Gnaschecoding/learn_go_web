package main

import (
	"context"
	"fmt"
)

type User struct {
	Name string
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "name", User{
		Name: "sb",
	})
	GetUserName(ctx)
}

func GetUserName(ctx context.Context) {
	fmt.Println(ctx.Value("name").(User))
}
