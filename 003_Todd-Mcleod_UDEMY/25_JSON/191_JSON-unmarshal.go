package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `[{"Nome":"Alex","Sobrenome":"Back","Idade":44,"Altura":1.75},{"Nome":"Norma","Sobrenome":"Sizilio","Idade":41,"Altura":1.55}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var pessoa []Shit

	err := json.Unmarshal(bs, &pessoa)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range pessoa {
		fmt.Println("\nPESSOA NÚMERO", i)
		fmt.Println(v.Nome, v.Sobrenome, v.Idade, v.Altura)
	}
}

type Shit struct {
	Nome      string  `json:"Nome"`
	Sobrenome string  `json:"Sobrenome"`
	Idade     int     `json:"Idade"`
	Altura    float64 `json:"Altura"`
}
