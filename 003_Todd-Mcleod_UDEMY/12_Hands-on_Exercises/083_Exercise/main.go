package main

import "fmt"

func main() {
	var i int
	for {
		fmt.Println("shit")
		i++
		if i >= 5 {
			fmt.Print("break")
			break
		}
	}
}

/*
● create a for loop that uses break statement
● increment or decrement the variable being checked as a condition in the loop
*/
