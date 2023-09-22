package main

import "fmt"

func media(lista []float64) float64 {
	var total float64
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() {
	minhaLista := []float64{2, 3, 5, 7, 11, 13}
	fmt.Println(media(minhaLista))
}
