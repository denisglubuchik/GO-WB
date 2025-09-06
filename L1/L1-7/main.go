package main

import (
	"math/rand"
	"sync"
)

type Map struct {
	mu   sync.Mutex
	data map[any]any
}

func NewMap() *Map {
	return &Map{mu: sync.Mutex{}, data: make(map[any]any)}
}

func (m *Map) Set(key any, value any) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	m := NewMap()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(i, rand.Intn(100))
		}(i)
	}
	wg.Wait()
}
