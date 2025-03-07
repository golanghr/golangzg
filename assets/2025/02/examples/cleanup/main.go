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
	// CAUTION: runtime.AddCleanup runs once the object
	// is no longer reachable, but not immediately
	// when the object is no longer reachable
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
	wp, ok := c.m[key]
	if !ok {
		return "", false
	}
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

	_, ok := cache.Get(1)
	fmt.Println("cached value OK:", ok)

	runtime.GC() // real work simulated

	_, ok = cache.Get(1)
	fmt.Println("cached value OK:", ok)
}
