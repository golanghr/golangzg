package main

import (
	"fmt"
	"runtime"
	"sync"
	"weak"
)

type Cache struct {
	m  map[int]weak.Pointer[string]
	mu sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		m: map[int]weak.Pointer[string]{},
	}
}

func (c *Cache) Add(key int, value *string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	wp := weak.Make(value)
	c.m[key] = wp
	runtime.AddCleanup(&value, func(key int) {
		c.mu.Lock()
		defer c.mu.Unlock()
		fmt.Println("value for key", key, "removed")
		delete(c.m, key)
	}, key)
}

func (c *Cache) Get(key int) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	wp := c.m[key]
	valPtr := wp.Value()
	if valPtr == nil {
		return "", false
	}
	return *valPtr, true
}

func main() {
    cache := NewCache()

    str := "Zagreb"
    cache.Add(1, &str)

    str = "Prague"
    cache.Add(1, &str)

    value, ok := cache.Get(1)
    fmt.Println("str:", str)
    fmt.Println("cached value OK:", ok, value)

    runtime.GC() // real work simulated

    value, ok = cache.Get(1)
    fmt.Println("cached value OK:", ok, value)

    runtime.GC()
}





