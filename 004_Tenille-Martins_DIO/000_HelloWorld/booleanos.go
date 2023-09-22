package main

import "fmt"

func main() {
	fmt.Println("true and true:", true && true)
	fmt.Println("true and false:", true && false)
	fmt.Println("false and false:", false && false)
	fmt.Println("true or true:", true || true)
	fmt.Println("true or false:", true || false)
	fmt.Println("false or false:", false || false)
	fmt.Println("not true:", !true)
	fmt.Println("not false:", !false)
}
