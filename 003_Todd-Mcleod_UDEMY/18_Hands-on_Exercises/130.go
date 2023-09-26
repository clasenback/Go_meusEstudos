package main

import "fmt"

func main() {
	ms := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Nome",
		friends:   map[string]int{"teta": 1, "perna": 3},
		favDrinks: []string{"teste1", "teste2", "teste3"},
	}

	fmt.Println(ms)
}

/*
Create and use an anonymous struct with these fields:
● first string
● friends map[string]int
● favDrinks []string
Print things.
*/
