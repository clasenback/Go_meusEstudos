package main

import "fmt"

func main() {
	x := "Alex"
	fmt.Printf("%v \t %T", &x, &x)
}

/*
● Create a value and assign it to a variable.
● Print the address of that value.
*/
