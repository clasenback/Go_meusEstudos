package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := "Alex1979"
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
	fmt.Println(bs)

	s := "Alex1979"
	err = bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil {
		fmt.Println(err)
		fmt.Println("LOGIN FALHOU")
		return
	}
	fmt.Println(len(bs))
	fmt.Println("LOGIN CONCLUÍDO")
}
