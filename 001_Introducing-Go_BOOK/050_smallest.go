package main

import "fmt"

func smallest(x []int) int {
	var sm int = x[0]
	for i := 0; i < len(x); i++ {
		if sm > x[i] {
			sm = x[i]
		}
	}
	return sm
}

func add(args ...int) (total int, it int) {
	for i, v := range args {
		total += v
		it = i + 1
	}
	return total, it
}

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(smallest(x))

	total, i := add(48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17)
	fmt.Println(total, "\t", i, "elementos")
}

/*
Write a program that finds the smallest number in this list:
x := []int{
48,96,86,68,
57,82,63,70,
37,34,83,27,
19,97, 9,17,
}
*/
