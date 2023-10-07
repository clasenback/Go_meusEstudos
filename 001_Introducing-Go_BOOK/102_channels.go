package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprint("pinger:", rand.Intn(1000))
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprint("ponger:", i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println("printer", msg)
		time.Sleep(time.Second * 1)
	}
}
