package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := fanIn(cronometro("Relógio 1", 3e3), cronometro("Relógio 2", 1e3))

	t := 1 * time.Second
	timeout := time.Tick(t)
	//var norma, alex int = 1

	tLoop := time.Now()
	for i := 1; i <= 10; i++ {
		select {
		case v := <-c:
			fmt.Printf("%v\t%v\t%v\n", i, time.Since(tLoop), v)
		case <-timeout:
			fmt.Printf("%v\t%v\ttimeout contou %v\n", i, time.Since(tLoop), t)
		}
	}
	fmt.Printf("Foram tomados 10 tempos aleatórios. Programa encerrado após %v.\n", time.Since(tLoop))
}

func cronometro(msg string, r int) <-chan string {
	c := make(chan string)
	go func() {
		var t time.Duration
		for i := 0; ; i++ {
			t = time.Duration(rand.Intn(r)) * time.Millisecond
			time.Sleep(t)
			c <- fmt.Sprintf("%s contou %v", msg, t)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string) //canal de saida
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*

O programa roda três cronômetros aleatórios em concorrência e imprime as 10 primeiras contagens.
O cronometro 'timeout' poderia estar dentro da função fanIn, mas foi deixada no select em func main por motivos pedagógicos.

code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
