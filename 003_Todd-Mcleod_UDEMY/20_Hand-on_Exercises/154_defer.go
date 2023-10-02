package main

import "fmt"

func main() {
	fmt.Println("LIFO:")
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	defer fmt.Println("Defer 4")
	fmt.Println("LAST IN FIRST OUT")
}

/*
● “defer” multiple functions in main
	○ show that a deferred func runs after the func containing it exits.
	○ determine the order in which the multiple defer funcs run
*/
