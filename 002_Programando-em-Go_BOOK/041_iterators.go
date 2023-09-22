package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 0
	for {
		n++
		i := rand.Intn(4200)
		fmt.Printf("%d \t %d \t %d \n", n, i, i%42)
		if i%42 == 0 {
			break
		}
	}
	fmt.Printf("Saída após %d iterações. \n", n)
}
