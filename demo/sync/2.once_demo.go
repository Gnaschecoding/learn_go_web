package sync

import (
	"fmt"
	"sync"
)

type Once struct {
	close sync.Once
}

// 这个地方如果不用指针的话，后面的每次调用都相当于被复制一次
func (o *Once) OnceClose() {
	o.close.Do(func() {
		fmt.Printf("hello\n")
	})
}
