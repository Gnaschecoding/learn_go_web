package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(T *testing.T) {
	pool := sync.Pool{
		//New: func() any {
		//	fmt.Println("hhh ,new")
		//	return []byte{}
		//},
		New: func() any {
			return &User{}
		},
	}
	u1 := pool.Get().(*User)
	u1.ID = 1
	u1.Name = "Tom"

	//如果不想要这个的话，放回去前要重置
	u1.Reset()

	pool.Put(u1)

	u2 := pool.Get().(*User)
	fmt.Println(u1, u2)

}

type User struct {
	ID   uint64
	Name string
}

func (u *User) Reset() {
	u.ID = 0
	u.Name = ""
}
