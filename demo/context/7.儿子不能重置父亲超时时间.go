package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bg := context.Background()
	Timectx, cancel := context.WithTimeout(bg, time.Second)

	Subcctx, cancel2 := context.WithTimeout(Timectx, 3*time.Second)

	go func() {
		select {
		case <-Subcctx.Done():
			err := Subcctx.Err()
			fmt.Println("timeout", err)
			return
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	cancel2()

}
