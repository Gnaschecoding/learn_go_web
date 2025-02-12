package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestOverride(t *testing.T) {
	sm := SafeMap[string, string]{
		values: make(map[string]string, 4),
	}

	go func() {
		time.Sleep(1 * time.Second)
		sm.LoadOrStoreV2("a", "b")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		sm.LoadOrStoreV2("a", "c")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		sm.LoadOrStoreV2("a", "d")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		sm.LoadOrStoreV2("a", "e")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello ")
}
