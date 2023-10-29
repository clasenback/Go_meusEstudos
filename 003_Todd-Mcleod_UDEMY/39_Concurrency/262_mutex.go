package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(3)
	go incrementor("Foo:")
	go incrementor("Bar:")
	go incrementor("Mil:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		for !mutex.TryLock() {
			fmt.Println("Estado do mutex", s, "esperando destravar.")
			time.Sleep(time.Duration(10) * time.Millisecond)
		}
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
