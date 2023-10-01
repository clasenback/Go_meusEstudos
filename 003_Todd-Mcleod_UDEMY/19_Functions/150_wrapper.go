package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := lerArquivo("140.txt")
	if err != nil {
		log.Fatalf("erro em main.go na leitura do arquivo: %s", err)
		os.Exit(1)
	}
	fmt.Println(xb)
	fmt.Printf("%T\n", xb)
	/* for _, c := range xb {
		fmt.Print(string(c))
	} */
	fmt.Println()
	fmt.Println(string(xb))

}

func lerArquivo(f string) ([]byte, error) {
	xb, err := os.ReadFile(f)
	//defer os.Close(f)
	if err != nil {
		return nil, fmt.Errorf("Houve um erro com o arquivo: %s", err)
	}
	return xb, nil
}

/* func cronometar(f func(string) ([]byte, error)) ([]byte, error) {
	inicio := time.Now()
	xb, err := f()
	fmt.Println("programa executado em", time.Since(inicio))
	return xb, err
}
*/
