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
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	wg.Add(1)
	go func() {
		ip, err := Getip1(ctx, &wg)
		fmt.Println(ip, err)
	}()
	wg.Wait()
	fmt.Println("运行时间为", time.Since(t1))
}

func Getip1(ctx context.Context, wg *sync.WaitGroup) (ip string, err error) {

	go func() {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			fmt.Println("暂停输出", err)
			wg.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.0.0.1"
	wg.Add(1)
	return
}
