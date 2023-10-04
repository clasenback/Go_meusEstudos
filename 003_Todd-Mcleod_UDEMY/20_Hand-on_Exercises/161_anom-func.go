package main

import "fmt"

func main() {
	f := func() bool {
		return true
	}
	if x := f(); x {
		fmt.Println("f():\t", f())
	}

	g := func() bool {
		return true
	}
	if g() {
		fmt.Println("!g():\t", !g())
	}

	h := func() {
		fmt.Println("Teste")
	}
	for i := 0; i < 1; i++ {
		h()
	}
}
