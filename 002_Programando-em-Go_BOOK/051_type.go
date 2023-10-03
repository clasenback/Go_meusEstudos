package main

import "fmt"

type listaCompras []string

func main() {
	lista := make(listaCompras, 6)
	lista[0] = "Alface"
	lista[1] = "Pepino"
	lista[2] = "Azeite"
	lista[3] = "Atum"
	lista[4] = "Frango"
	lista[5] = "Chocolate"
	fmt.Println(lista)

	vegetais, carnes, outros := lista.Categorizar()
	fmt.Println("Vegetais:", vegetais)
	fmt.Println("Carnes:", carnes)
	fmt.Println("Outros:", outros)

	// Conversão de tipos
	extras := []string{"detergente", "vinho"}
	lista = append(lista, listaCompras(extras)...)
	_, _, outros = lista.Categorizar()
	fmt.Println("Outros:", outros)

}

func (lista listaCompras) Categorizar() ([]string, []string, []string) {
	var vegetais, carnes, outros []string
	for _, e := range lista {
		switch e {
		case "Alface", "Pepino":
			vegetais = append(vegetais, e)
		case "Atum", "Frango":
			carnes = append(carnes, e)
		default:
			outros = append(outros, e)
		}
	}
	return vegetais, carnes, outros
}
