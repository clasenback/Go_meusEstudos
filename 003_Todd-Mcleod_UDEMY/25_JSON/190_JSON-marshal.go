package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	eu := Pessoa{"Alex", "Back", 44, 1.75}
	norma := Pessoa{"Norma", "Sizilio", 41, 1.55}
	pessoas := []Pessoa{eu, norma}
	fmt.Println(pessoas)

	j, err := json.Marshal(pessoas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(j)
	fmt.Println(string(j))

}

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Altura    float32
}
