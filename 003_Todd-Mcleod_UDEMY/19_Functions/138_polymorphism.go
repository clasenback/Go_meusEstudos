package main

import "fmt"

func main() {
	livro := book{"Atlas", 104}
	var conta contagem = 59
	fmt.Println(livro)
	fmt.Println(conta)
	fmt.Printf("%#v \t%T \n", livro)
	fmt.Printf("%#v \t%T \n", conta)
}

type book struct {
	nome    string
	paginas int
}

func (b book) String() string {
	return fmt.Sprint("===", b.nome, b.paginas, "páginas ===")
}

type contagem int

func (c contagem) String() string {
	return fmt.Sprint("+++", c, "+++")
}
