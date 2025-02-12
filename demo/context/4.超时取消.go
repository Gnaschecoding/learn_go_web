package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	t1 := time.Now()
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	go func() {
		ip, err := Getip2(ctx, &wg)
		fmt.Println("ip:", ip, "err:", err)
	}()
	wg.Wait()
	fmt.Println("用时", time.Since(t1))
}

func Getip2(ctx context.Context, wg *sync.WaitGroup) (ip string, err error) {

	go func() {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			fmt.Println("取消传递", err)
			wg.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.0.0.1"
	wg.Done()
	return
}
