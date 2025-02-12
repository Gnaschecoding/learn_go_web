package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	childctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	end := make(chan struct{}, 1)
	go func() {
		MyBusiness()
		end <- struct{}{}
	}()
	ch := childctx.Done()

	select {
	case <-ch:
		fmt.Println("超时")
	case <-end:
		fmt.Println("正常执行")
	}
	
}

func MyBusiness() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Hello World!")
}
