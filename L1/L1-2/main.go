package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println("arr:", arr)

	for _, num := range arr {
		go func(num int) {
			fmt.Println(num, "^2 =", num*num)
		}(num)
	}
	time.Sleep(time.Second)
}
