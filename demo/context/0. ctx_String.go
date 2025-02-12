package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)
	b := []int{1, 2, 3, 4}
	fmt.Println(cap(b))
	b = append(b, 5, 6)
	fmt.Println(cap(b))
	b = append(b, []int{7, 8}...)
	fmt.Println(b)
	ctx.Err()
}

type Cache interface {
	Get(string) (string, error)
}
type OtherCache interface {
	GetValue(ctx context.Context, key string) (string, error)
}
type CacheAdapter struct {
	Cache
}

func (ca *CacheAdapter) GetValue(ctx context.Context, key string) (string, error) {
	return ca.Cache.Get(key)
}

type memoryMap struct {
	m map[string]string
}

func (m *memoryMap) Get(s string) (string, error) {
	return m.m[s], nil
}

type SafeCache struct {
	Cache
	sync.RWMutex
}

func (c *SafeCache) Get(key string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	return c.Cache.Get(key)
}

var c = &SafeCache{
	Cache: &memoryMap{},
}
