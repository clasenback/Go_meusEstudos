package main

import "fmt"

func main() {
	capitais := map[string]string{
		"GO": "Goiania",
		"PB": "João Pessoa",
		"PR": "Curitiba",
	}
	capitais["RN"] = "Natal"
	capitais["AM"] = "Manaus"
	capitais["SE"] = "Aracaju"

	populacao := make(map[string]int, len(capitais))
	pops := []int{6434052, 3914418, 10997462, 3373960, 3807923, 2228489}
	var i int
	for k, _ := range capitais {
		populacao[k] = pops[i]
		i++
	}

	fmt.Println(capitais)
	fmt.Println(populacao)

	estados := make(map[string]Estado, 6)
	estadosNomes := []string{"Goiás", "Paraíba", "Paraná", "Rio Grande do Norte", "Amazonas", "Sergipe"}

	i = 0
	for est, capital := range capitais {
		estados[est] = Estado{estadosNomes[i], pops[i], capital}
		i++
	}

	fmt.Println(estados)

	for sigla, estado := range estados {
		fmt.Printf("%s (%s) possui %d habitantes e sua capital é %s. \n", estado.nome, sigla, estado.populacao, estado.capital)
	}

	// as populações normalmente saem erradas, pois os maps não são ordenados
	// logo, o map estados é populado fora de ordem.
}

type Estado struct {
	nome      string
	populacao int
	capital   string
}
