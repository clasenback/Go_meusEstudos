package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Create("140.txt")
	if err != nil {
		log.Fatal("erro %s", err)
	}
	defer f.Close()

	s := []byte("Olá mundo!\n")
	p := pessoa{nome: "Alex", idade: 44}
	pp := []byte(p.String())

	var n int

	n, err = f.Write(s)
	fmt.Println(n, "caracteres escritos no arquivo.")
	if err != nil {
		log.Fatalf("Erro %s", err)
	}

	n, err = f.Write(pp)
	fmt.Println(n, "caracteres escritos no arquivo.")
	if err != nil {
		log.Fatalf("Erro %s", err)
	}

	n, err = p.WriteOut(f)
	fmt.Println(n, "caracteres escritos no arquivo.")
	if err != nil {
		log.Fatalf("Erro %s", err)
	}

}

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) WriteOut(w io.Writer) (n int, err error) {
	n, err = w.Write([]byte(p.nome))
	return n, err
}

func (p pessoa) String() string {
	return fmt.Sprintf("%s tem %v anos.\n", p.nome, p.idade)
}
