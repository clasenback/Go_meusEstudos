package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"usuarios"`
}

type User struct {
	Nome   string `json:"Nome"`
	Tipo   string `json:"Tipo"`
	Idade  int    `json:"Idade"`
	social Social `json:"rede social"`
}

type Social struct {
	facebook string `json:"facebook"`
	twitter  string `json:"twitter"`
}

func main() {

	jsonFile, err := os.Open("usuarios.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Arquivo aberto com sucesso")

	byteValue, _ := io.ReadAll(jsonFile)

	var usuarios Users

	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.Users); i++ {
		fmt.Println("--------------------------")
		fmt.Println("Usuario Tipo: " + usuarios.Users[i].Tipo)
		fmt.Println("Usuário Idade: " + strconv.Itoa(usuarios.Users[i].Idade))
		fmt.Println("Usuário Nome: " + usuarios.Users[i].Nome)
	}
}
