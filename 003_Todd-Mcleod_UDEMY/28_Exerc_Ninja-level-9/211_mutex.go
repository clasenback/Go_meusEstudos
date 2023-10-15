package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	var inc int
	var wg sync.WaitGroup
	var cadeado sync.Mutex
	ii := func() func() int {
		iii := 0
		return func() int {
			iii++
			return iii
		}
	}()
	loopFor := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			t := time.Now()
			cadeado.Lock()
			gs := ii()
			fmt.Println("Goroutine", gs, "iniciada")
			v := inc
			v++
			dorme := time.Duration(rand.Intn(2)) * time.Millisecond
			// time.Sleep(dorme)
			inc = v
			fmt.Println("Goroutine", gs, "concluída em", time.Since(t), "após dormir", dorme)
			wg.Done()
			cadeado.Unlock()
		}()
		fmt.Println("Goroutines ativas:", runtime.NumGoroutine()-1)
	}
	fmt.Println("Loop For concluído em", time.Since(loopFor))
	wg.Wait()
	fmt.Println("Programa concluído. Inc =", inc)
}

/*
Fix the race condition you created in the previous exercise by using a mutex
● it makes sense to remove runtime.Gosched()
*/
