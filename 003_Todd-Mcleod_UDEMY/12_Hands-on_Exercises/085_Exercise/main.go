package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%v \t %v \n", i, j)
		}
	}
}

/*
● create a loop that runs 5 times
● nest a loop within the first loop that also prints 5 times
● print something in each loop to illustrate what is occuring
*/
