package main

import (
	"fmt"
	"os"
)

func main() {
	tmp := os.TempDir()

	CriarArquivos(tmp)
	CriarArquivos(tmp, "teste1")
	CriarArquivos(tmp, "teste2", "teste3", "teste4")
	CriarArquivos(tmp, "teste1")
}

func CriarArquivos(dirBase string, arquivos ...string) {
	for _, nome := range arquivos {
		caminhoArq := fmt.Sprintf("%s/%s.%s", dirBase, nome, "txt")
		arq, err := os.Create(caminhoArq)
		defer arq.Close()
		if err != nil {
			fmt.Println("Erro ao criar arquivo %s: %v\n", nome, err)
			os.Exit(1)
		}
		fmt.Printf("Arquivo %s criado.\n", arq.Name())
	}
}
