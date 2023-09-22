package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i > 20 {
			break
		}
		if i%3 > 0 {
			fmt.Println(i)
			continue
		} else {
			fmt.Println(i, "é múltiplo de 3.")
		}

	}
}
