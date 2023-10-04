package main

import "fmt"

func main() {
	f := proxQuadrado()
	for i := 1; i <= 10; i++ {
		fmt.Println(i, f())
	}
}

func proxQuadrado() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}
