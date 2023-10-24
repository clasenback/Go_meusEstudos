package main

import "fmt"

func main() {
	shit := customErr{erro: "caguei o programa"}
	foo(shit)
}

type customErr struct {
	erro string
}

func (err customErr) Error() string {
	return fmt.Sprint("Meu erro: ", err.erro)
}

func foo(err error) {
	fmt.Println("dentro de foo:", err)
}

/*
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”.
*/
