package main

import "fmt"

type Arquivo struct {
	nome       string
	tamanho    float64
	caracteres int
	palavras   int
	linhas     int
}

func main() {
	arquivo := Arquivo{"artigo.txt", 12.68, 12_986, 1_862, 220}
	fmt.Printf("%#v\n", arquivo)

	codigoFonte := Arquivo{tamanho: 1.12, nome: "programa.go"}
	fmt.Printf("%#v\n", codigoFonte)

	fmt.Printf("%s\t%.2fKB\n", arquivo.nome, arquivo.tamanho)
	fmt.Printf("%s\t%.2fKB\n", codigoFonte.nome, codigoFonte.tamanho)

	ptrArquivo := &Arquivo{tamanho: 7.29, nome: "arquivo.txt"}
	fmt.Printf("%s\t%.2fKB\n", ptrArquivo.nome, ptrArquivo.tamanho)
	fmt.Printf("%v\n", ptrArquivo)

	fmt.Printf("Média de palavras por linha: \t%.2f\n", arquivo.MediaDePalavrasPorLinha())
	fmt.Printf("Tamanho médio de palavra: \t%.2f\n", arquivo.TamanhoMedioDePalavra())
}

func (arq *Arquivo) TamanhoMedioDePalavra() float64 {
	return float64(arq.caracteres) / float64(arq.palavras)
}

func (arq *Arquivo) MediaDePalavrasPorLinha() float64 {
	return float64(arq.palavras) / float64(arq.linhas)
}
