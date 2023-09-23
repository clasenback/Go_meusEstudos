package main

import "fmt"

var idade int = 44

const nome string = "Alexander"

func main() {
	apelido := nome[3:8]
	fmt.Printf("Seu nome é %v, seu apelido é %v e sua idade atual é %v.", nome, apelido, idade)
}

/*
● create the following variables with the following scopes:
	○ package level
		■ create outside of func main
		■ use the
			● var keyword
			● const keyword
	○ block level
		■ inside func main
		■ use the short declaration operator
			● use the variables in func mai
*/
