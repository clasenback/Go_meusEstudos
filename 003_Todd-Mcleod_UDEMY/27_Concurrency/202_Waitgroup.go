package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU(), "CPUs")
	fmt.Println(runtime.NumGoroutine(), "gouroutines")
	go funcao("shit")
	wg.Add(1)
	fmt.Println(runtime.NumGoroutine(), "gouroutines")
	//wg.Add(1)
	go funcao("cum")
	wg.Add(1)
	fmt.Println(runtime.NumGoroutine(), "gouroutines")
	wg.Wait()
	fmt.Println(runtime.NumGoroutine(), "gouroutines")
}

func funcao(s string) {
	t := time.Now()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println(s, "concluiu sua execução em", time.Since(t))
	wg.Done()
}

/* wg.Add(1)
fmt.Println(wg, runtime.NumGoroutine())
wg.Add(1)
fmt.Println(wg, runtime.NumGoroutine())
wg.Add(1)
fmt.Println(wg, runtime.NumGoroutine()) */
/* 	fmt.Println(wg, runtime.NumGoroutine())
wg.Done()
fmt.Println(wg, runtime.NumGoroutine())
wg.Done()
fmt.Println(wg, runtime.NumGoroutine())
wg.Done()
fmt.Println(wg, runtime.NumGoroutine()) */
