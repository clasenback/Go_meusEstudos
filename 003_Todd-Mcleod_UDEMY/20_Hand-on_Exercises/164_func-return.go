package main

import "fmt"

func main() {
	f := f42()
	fmt.Println("f:\t", f)
	fmt.Println("f():\t", f())
}

func f42() func() int {
	return func() int {
		return 42
	}
}

/*
● Create a func
	○ which returns a func
		■ which returns 42
● assign the returned func to a variable
● call the returned func
● print
*/
