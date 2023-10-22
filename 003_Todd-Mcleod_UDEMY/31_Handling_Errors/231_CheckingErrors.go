package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var filename string
	fmt.Print("Digite o nome do arquivo: ")
	_, err := fmt.Scan(&filename)
	checkError(err)
	fmt.Println(filename)

	f, err := os.Create(filename)
	checkError(err)
	defer f.Close()
	fmt.Println("Arquivo", filename, "criado.")

	var texto string
	fmt.Println("Digite o texto para o arquivo: ")
	n, err := fmt.Scan(&texto)
	r := strings.NewReader(texto)
	checkError(err)
	fmt.Println(n, "bytes digitados.")

	nn, err := io.Copy(f, r)
	checkError(err)
	fmt.Println(nn, "bytes gravados.")

	f, err = os.Open(filename)
	checkError(err)
	defer f.Close()

	bs, err := io.ReadAll(f)
	checkError(err)
	fmt.Println(string(bs))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("erro recebido")
		panic(err)
	}
}
