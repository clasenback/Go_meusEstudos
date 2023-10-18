package main

import "fmt"

func main() {
	c := contador()
	for i := 0; i < 10; i++ {
		fmt.Println(c())
	}
}

func contador() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
