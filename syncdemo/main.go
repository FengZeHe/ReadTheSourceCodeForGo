package main

import (
	"log"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	count := &Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Inc()
		}()
	}
	wg.Wait()
	log.Println("count:", count.Value())
}
