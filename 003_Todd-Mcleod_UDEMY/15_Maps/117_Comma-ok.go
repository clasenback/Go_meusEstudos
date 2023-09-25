package main

import "fmt"

func main() {
	am := make(map[string]string)
	am["Alex"] = "Back"
	am["Norma"] = "Sizilio"
	am["Lourdes"] = "Clasen"
	am["Juliane"] = "Nacari"

	keys := []string{"Alex", "Norma", "Lourdes", "Juliane", "Gervásio"}

	for _, k := range keys {
		if v, ok := am[k]; !ok {
			fmt.Println("chave", k, "não existe")
		} else {
			fmt.Println(k, v)
		}
	}
}
