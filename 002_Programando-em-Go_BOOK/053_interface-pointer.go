package main

import "fmt"

type listaGenerica []interface{}

func main() {
	lista := listaGenerica{1, "Café", 42, true, 23, "Bola", 3.14, false}

	fmt.Printf("Lista original:\t %v \n", lista)
	fmt.Printf("Removendo do início: \t %v,\n após remoção: \t %v \n", lista.RemoverInicio(), lista)
	fmt.Printf("Removendo do fim: \t %v,\n após remoção: \t %v \n", lista.RemoverFim(), lista)
	fmt.Printf("Removendo do índice 3: \t %v,\n após remoção: \t %v \n", lista.RemoverIndice(3), lista)
	fmt.Printf("Removendo do índice 0: \t %v,\n após remoção: \t %v \n", lista.RemoverIndice(0), lista)
	fmt.Printf("Removendo do último índice: \t %v,\n após remoção: \t %v \n", lista.RemoverIndice(len(lista)-1), lista)
}

func (lista *listaGenerica) RemoverIndice(indice int) interface{} {
	l := *lista
	removido := l[indice]
	*lista = append(l[0:indice], l[indice+1:]...)
	return removido
}

func (lista *listaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

func (lista *listaGenerica) RemoverFim() interface{} {
	return lista.RemoverIndice(len(*lista) - 1)
}
