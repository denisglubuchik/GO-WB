package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 5
	timeout := time.After(time.Duration(n) * time.Second)
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(100)
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case <-timeout:
			fmt.Println("timeout")
			return
		case num := <-ch:
			fmt.Println("Fetched", num)
		}

	}
}
