package main

import "fmt"

func main() {
	s := [5]int{2, 3, 5, 7, 11}
	for i, v := range s {
		fmt.Println(i, v)
	}

	m := map[string]int{
		"dois":  2,
		"três":  3,
		"cinco": 5,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
