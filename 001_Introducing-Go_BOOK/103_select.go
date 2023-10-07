package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go from1(c1)
	go from2(c2)
	go seleciona(c1, c2)

	var input string
	fmt.Scanln(&input)
}

func from1(c1 chan string) {
	t0 := time.Now()
	for {
		time.Sleep(time.Second * 2)
		c1 <- fmt.Sprint("from 1: ", time.Since(t0))
	}
}

func from2(c2 chan string) {
	t0 := time.Now()
	for {
		time.Sleep(time.Second * 3)
		c2 <- fmt.Sprint("from 2: ", time.Since(t0))
	}
}

func seleciona(c1, c2 chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(time.Second):
			fmt.Println("timeout")
			/* default:
			fmt.Println("nothing ready") */
		}
	}
}
