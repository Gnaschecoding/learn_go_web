package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	childctx := context.WithValue(ctx, "key1", "valu1")
	cchildctx := context.WithValue(childctx, "key2", "valu2")

	//父亲的节点拿不到后面子孙的值，但是子孙的可以获取父亲的值
	val1 := childctx.Value("key2")
	fmt.Println("val1:", val1)
	val2 := cchildctx.Value("key1")
	fmt.Println("val2:", val2)

}
