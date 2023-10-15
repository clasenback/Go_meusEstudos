package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			t := rand.Intn(100)
			time.Sleep(time.Duration(t) * time.Millisecond)
			//runtime.Gosched()
			v++
			fmt.Println("i:", i, "\tcount:", v, "\tativas:", runtime.NumGoroutine(), "\trand:", t)
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
