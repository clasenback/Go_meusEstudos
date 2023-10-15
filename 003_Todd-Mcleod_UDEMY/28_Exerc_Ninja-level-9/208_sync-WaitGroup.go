package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	incrementa := incrementador()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go rotina(time.Duration(rand.Intn(500))*time.Millisecond, incrementa(), &wg)
	}
	wg.Wait()
}

func incrementador() func() int {
	j := 0
	return func() int {
		j++
		return j
	}
}

func rotina(t time.Duration, i int, wg *sync.WaitGroup) {
	tIni := time.Now()
	time.Sleep(t)
	fmt.Println("Goroutine", i, "concluida em", time.Since(tIni))
	wg.Done()
}

/*
● in addition to the main goroutine, launch two additional goroutines
	○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/
