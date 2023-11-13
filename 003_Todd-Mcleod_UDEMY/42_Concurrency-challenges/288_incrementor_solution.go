package main

import (
	"fmt"
)

var count int

//var wg sync.WaitGroup

func main() {
	//wg.Add(2)
	out12 := make(chan int)
	out21 := make(chan int)
	incrementor("1", out12, out21)
	incrementor("2", out21, out12)
	out12 <- 0
	//wg.Wait()

	for v := range out {
		fmt.Println("Final Counters:", count)
	}

}

func incrementor(s string, out chan<- int, in <-chan int) {
	go func() {
		for i := 0; i < 20; i++ {
			count = <-in
			//atomic.AddInt64(&count, 1)
			fmt.Println("Process:", s, "printing:", i)
		}
	}()

	//wg.Done()
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
