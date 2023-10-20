package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type pacote struct {
	n     int
	dados int
}

func main() {
	fmt.Println("Programa iniciado:", runtime.NumGoroutine(), "gouroutines")

	c := make(chan pacote)
	quit := make(chan bool)
	var pct pacote
	var wg sync.WaitGroup

	wg.Add(1)
	go func(c chan<- pacote, q chan<- bool) {
		for i := 0; i < 10; i++ {
			pct.n = i
			pct.dados = rand.Intn(1000)
			// time.Sleep(time.Duration(pct.dados) * time.Microsecond)
			c <- pct
		}
		close(c)
		q <- true
		close(q)
		wg.Done()
	}(c, quit)
	fmt.Println("Envio de dados iniciado:", runtime.NumGoroutine(), "gouroutines")

	wg.Add(1)
	go func(c <-chan pacote) {
		for {
			v, ok := <-c
			if ok == false {
				break
			}
			fmt.Println(v.n, v.dados, "\tloop for comma, ok")
		}
		wg.Done()
	}(c)
	fmt.Println("Loop for comma ok iniciado:", runtime.NumGoroutine(), "gouroutines")

	wg.Add(1)
	go func(c <-chan pacote) {
		for v := range c {
			fmt.Println(v.n, v.dados, "\tloop for range")
		}
		wg.Done()
	}(c)
	fmt.Println("Loop for range iniciado:", runtime.NumGoroutine(), "gouroutines")

	wg.Add(1)
	go func(c <-chan pacote, q <-chan bool) {
		for {
			select {
			case v := <-c:
				fmt.Println(v.n, v.dados, "\tloop select")
			case <-quit:
				fmt.Println("quit select")
				wg.Done()
				return
			}
		}
	}(c, quit)
	fmt.Println("Loop for select iniciado:", runtime.NumGoroutine(), "gouroutines")

	wg.Wait()
	fmt.Println("Programa concluido:", runtime.NumGoroutine(), "gouroutines")

}

/*
● write a program that
○ puts 100 numbers to a channel
○ pull the numbers off the channel and print them
*/
