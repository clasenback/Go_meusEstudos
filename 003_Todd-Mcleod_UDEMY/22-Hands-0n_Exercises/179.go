package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walk.")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

/*
In the provided code, correct the code to do the following:
	● implement the "youngin" interface
		○ you can change anything in the code to do so except the interface
	● maintain consistency within the code
		○ receivers should stick with either POINTER or VALUE semantics
	● when choosing between POINTER or VALUE semantics, understand why you chose one or the other
*/
