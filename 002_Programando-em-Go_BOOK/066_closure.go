package main

import "fmt"

func main() {
	f := fib()
	g := seq()
	for i := 0; i <= 10; i++ {
		fmt.Println(g(), f())
	}

}

func seq() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
