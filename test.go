package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v 	map[string] int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	for i := 0; i < 500; i++ {
		go c.Inc("222")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
	fmt.Println(c.Value("222"))
	fmt.Println(c)
}

