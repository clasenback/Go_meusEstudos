package main

import (
	"fmt"
	"math"
)

func main() {
	quadrado := SQUARE{length: 100, width: 40}
	circulo := CIRCLE{radius: 80}

	INFO(quadrado)
	INFO(circulo)
}

type SQUARE struct {
	length float64
	width  float64
}

type CIRCLE struct {
	radius float64
}

func (s SQUARE) AREA() float64 {
	return s.length * s.width
}

func (c CIRCLE) AREA() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type SHAPE interface {
	AREA() float64
}

func INFO(s SHAPE) {
	fmt.Println(s.AREA())
}

/*
● create a type SQUARE
	○ length float64
	○ width float64
● create a type CIRCLE
	○ radius float64
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
		■ math.Pi
		■ math.Pow
	○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/
