package main

import (
	"log"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("k1", 1)
	m.Store("k2", 2)

	if k, ok := m.Load("k1"); ok {
		log.Println(k)
	}

	c, loaded := m.LoadOrStore("k3", 3)
	log.Println(c, loaded)

	m.Delete("k3")
	m.Range(func(k, v interface{}) bool {
		log.Println(k, v)
		return true
	})
}
