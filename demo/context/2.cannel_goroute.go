package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var Wait = sync.WaitGroup{}

func main() {
	t1 := time.Now()
	ctx, cannel := context.WithCancel(context.Background())

	Wait.Add(1)
	go func() {
		ip, err := Getip(ctx)
		fmt.Println("ip:", ip, "err:", err)

	}()

	go func() {
		time.Sleep(2 * time.Second)
		cannel()
	}()

	Wait.Wait()
	fmt.Println("运行时间为：", time.Since(t1))
}

func Getip(ctx context.Context) (ip string, err error) {

	go func() {
		select {
		case <-ctx.Done():

			err = ctx.Err()
			fmt.Println("输出取消", err)
			Wait.Done()
			return
		}
	}()

	time.Sleep(4 * time.Second)
	ip = "192.0.0.1"
	Wait.Done()
	return
}
