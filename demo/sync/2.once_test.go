package sync

import (
	"fmt"
	"testing"
)

func TestOnce(t *testing.T) {
	once := &Once{}
	once.OnceClose()
	once.OnceClose()
	fmt.Printf("over")
}
