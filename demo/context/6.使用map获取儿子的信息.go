package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	childctx := context.WithValue(ctx, "map", map[string]string{})

	cchildctx := context.WithValue(childctx, "key2", "val2")

	m := cchildctx.Value("map").(map[string]string)
	m["key2"] = cchildctx.Value("key2").(string)
	val1 := childctx.Value("key2")
	fmt.Println("val1:", val1)
	val2 := childctx.Value("map")
	fmt.Println("val2:", val2)
}
