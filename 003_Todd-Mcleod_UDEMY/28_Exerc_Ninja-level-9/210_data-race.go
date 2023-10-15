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
	ii := func() func() int {
		iii := 0
		return func() int {
			iii++
			return iii
		}
	}()
	loopFor := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			t := time.Now()
			gs := ii()
			fmt.Println("Goroutine", gs, "iniciada")
			v := inc
			v++
			dorme := time.Duration(rand.Intn(2)) * time.Millisecond
			time.Sleep(dorme)
			inc = v
			fmt.Println("Goroutine", gs, "concluída em", time.Since(t), "após dormir", dorme)
			wg.Done()
		}()
		fmt.Println("Goroutines ativas:", runtime.NumGoroutine()-1)
	}
	fmt.Println("Loop For concluído em", time.Since(loopFor))
	wg.Wait()
	fmt.Println("Programa concluído. Inc =", inc)
}

/*
● Using goroutines, create an incrementer program
	○ have a variable to hold the incrementer value
	○ launch a bunch of goroutines
		■ each goroutine should
			● read the incrementer value
				○ store it in a new variable
			● yield the processor with runtime.Gosched()
			● increment the new variable
			● write the value in the new variable back to the incrementer variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag
● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
*/
