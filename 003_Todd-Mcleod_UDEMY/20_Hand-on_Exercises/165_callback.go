package main

import "fmt"

func main() {
	dobro := func(n int) (string, int) {
		return "dobro", 2 * n
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(printResposta(dobro, i), "\t", printResposta(square, i))
	}

}

func square(n int) (string, int) {
	return "quadrado", n * n
}

//ao criar e empregar o tipo callBack,
//melhor-se a legibilidade da função que toma callback como parametro
//callBack é mais fácil de entender que 'func(int) (string, int)'
type callBack func(int) (string, int)

func printResposta(f callBack, n int) string {
	s, v := f(n)
	return fmt.Sprintf("O %s de %v é %v", s, n, v)
}

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
	● pass a func into a func as an argument
		○ func square(n int) int
		○ printSquare(f func(int)int, int) string
*/
