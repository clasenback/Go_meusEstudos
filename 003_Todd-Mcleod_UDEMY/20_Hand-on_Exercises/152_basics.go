package main

import "fmt"

func main() {
	fmt.Println(foo())
	x, y := bar()
	fmt.Println(x, y)

}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "bar"
}

/*
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
*/
