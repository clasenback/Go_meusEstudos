package main

import "fmt"

func main() {
	x := foo()
	y := bar()

	fmt.Printf("foo: \t%v \t %T\n", foo, foo)
	fmt.Printf("x:   \t%v \t\t %T\n", x, x)
	fmt.Printf("bar: \t%v \t %T\n", bar, bar)
	fmt.Printf("y:   \t%v \t %T\n", y, y)
	fmt.Printf("y(): \t%v \t\t %T\n", y(), y())
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
