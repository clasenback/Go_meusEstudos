package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var total int
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println(i, "\t", x)
			total++
		}
	}
	fmt.Println(total)
}

/*
● use the "statement statement" idiom to
	○ initialize x with and random int between 0 inclusive and 5 exclusive
	○ if x is 3
		■ print "x is 3"
● run that code 100 times
*/
