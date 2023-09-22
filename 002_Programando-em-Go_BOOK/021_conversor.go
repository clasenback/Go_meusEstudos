package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		fmt.Println("As opções válidas para unidades são:")
		fmt.Println("C \t Celsius")
		fmt.Println("F \t Fahrenheit")
		fmt.Println("km \t quilômetros")
		fmt.Println("mi \t milhas")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	for _, valorOrigem := range valoresOrigem {
		converter(valorOrigem, unidadeOrigem)
	}
}

func converter(valorOrigem string, unidadeOrigem string) {
	v, err := strconv.ParseFloat(valorOrigem, 64)
	if err != nil {
		fmt.Printf("O valor %s não é uma número válido! \n", valorOrigem)
	}
	var unidadeDestino string
	var valorDestino float64
	switch unidadeOrigem {
	case "C":
		unidadeDestino = "F"
		valorDestino = v*1.8 + 32
		fmt.Printf("%.2f %s = %.2f %s \n", v, unidadeOrigem, valorDestino, unidadeDestino)
	case "F":
		unidadeDestino = "C"
		valorDestino = (v - 32) / 1.8
		fmt.Printf("%.2f %s = %.2f %s  \n", v, unidadeOrigem, valorDestino, unidadeDestino)
	case "km":
		unidadeDestino = "mi"
		valorDestino = v / 1.60934
		fmt.Printf("%.2f %s = %.2f %s \n", v, unidadeOrigem, valorDestino, unidadeDestino)
	case "mi":
		unidadeDestino = "km"
		valorDestino = v * 1.60934
		fmt.Printf("%.2f %s = %.2f %s \n", v, unidadeOrigem, valorDestino, unidadeDestino)
	default:
		fmt.Printf("%v não é uma unidade conhecida!", unidadeOrigem)
		os.Exit(1)
	}
}
