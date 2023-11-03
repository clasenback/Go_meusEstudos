package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	c1 := make(chan pacote)
	c2 := make(chan pacote)

	//gera 10 valores e envia pelo canal c1
	go populate(c1)

	//recebe os valores de c1 processa e envia o resultado pelo canal c2
	go fanOutIn(c1, c2)

	//recebe os valores de c2 e imprime
	for pct := range c2 {
		fmt.Printf("pacote %v processado: %v em %v\n", pct.nro, pct.conteudo, pct.tempo)
	}

	fmt.Println("programa concluído em", time.Since(t))
}

type pacote struct {
	nro      int
	conteudo int
	tempo    time.Duration
}

func populate(c chan pacote) {
	var pct pacote
	for i := 0; i < 10; i++ {
		pct.nro = i
		pct.conteudo = i
		fmt.Printf("pacote %v criado: %v em %v\n", pct.nro, pct.conteudo, pct.tempo)
		c <- pct
	}
	close(c)
}

func fanOutIn(c1, c2 chan pacote) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 pacote) {
			c2 <- doShit(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func doShit(pct pacote) pacote {
	pct.tempo = time.Millisecond * time.Duration(rand.Intn(500))
	//time.Sleep(pct.tempo)
	time.Sleep(time.Millisecond)
	pct.conteudo = pct.conteudo + rand.Intn(1000)
	return pct
}
