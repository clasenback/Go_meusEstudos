package main

import (
	"Go_meusEstudos/003_Todd-Mcleod_UDEMY/36_Ninja-level-12/255/quote"
	"Go_meusEstudos/003_Todd-Mcleod_UDEMY/36_Ninja-level-12/255/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	fmt.Println(word.UseCount("quilo gordo carro"))
}
