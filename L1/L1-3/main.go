package main

import (
	"fmt"
	"sync"
)

func create_worker(id int, c chan int, wg *sync.WaitGroup) {
	for message := range c {
		fmt.Println(message, "fetched by worker", id)
	}
	wg.Done()
}

func main() {
	n := 5
	var wg sync.WaitGroup

	ch := make(chan int, n)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go create_worker(i, ch, &wg)
	}
	for i := 0; i < n*10; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
