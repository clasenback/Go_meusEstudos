package main

import "fmt"

func main() {
	var minhasNotas [5]float64
	minhasNotas[0] = 7.9
	minhasNotas[1] = 6.6
	minhasNotas[2] = 8.5
	minhasNotas[3] = 9.3
	minhasNotas[4] = 5.75

	var total float64
	for i := 0; i < len(minhasNotas); i++ {
		total += minhasNotas[i]
	}

	fmt.Println(total / float64(len(minhasNotas)))
}
