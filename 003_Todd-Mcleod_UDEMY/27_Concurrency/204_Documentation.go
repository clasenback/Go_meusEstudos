package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- doShit(rand.Intn(5))
	}()
	fmt.Println(<-ch)
}

func doShit(n int) int {
	return n * n
}
