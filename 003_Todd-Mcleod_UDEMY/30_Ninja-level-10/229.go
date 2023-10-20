package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type pacote struct {
	gr       int
	nro      int
	conteudo int
}

func main() {
	fmt.Println(runtime.NumGoroutine(), "goroutines rodando")
	c := make(chan pacote)
	go fanOut(c, 10, 10)
	fmt.Println(runtime.NumGoroutine(), "goroutines rodando")
	fanIn(c)
	fmt.Println("Programa concluido")
}

func fanOut(c chan<- pacote, ngr int, npct int) {
	fmt.Println("Iniciando geração de", ngr*npct, "pacotes")
	var wg sync.WaitGroup
	for i := 0; i < ngr; i++ {
		wg.Add(1)
		go func(cc chan<- pacote, i int, npct int) {
			fmt.Println("Goroutine fanOut", i, "iniciada")
			var pct pacote
			for j := 0; j < npct; j++ {
				pct.conteudo = rand.Intn(500)
				time.Sleep(time.Duration(pct.conteudo) * time.Millisecond)
				pct.gr = i
				pct.nro = j
				cc <- pct
			}
			fmt.Println("Goroutine fanOut", i, "concluída")
			wg.Done()
		}(c, i, npct)
		fmt.Println(runtime.NumGoroutine(), "goroutines rodando (dentro de fanOut)")
	}
	wg.Wait()
	close(c)
}

func fanIn(c <-chan pacote) {
	soma := 0
	for v := range c {
		fmt.Println("goroutine", v.gr, "\tpacote", v.nro, "recebido:\t", v.conteudo)
		soma++
	}
	fmt.Println(soma, "pacotes recebidos")
}

/*
● write a program that
	○ launches 10 goroutines
		■ each goroutine adds 10 numbers to a channel
	○ pull the numbers off the channel and print them
*/
