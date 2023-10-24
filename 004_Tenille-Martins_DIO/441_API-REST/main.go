package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Crie um roteador usando o Gorilla Mux
	r := mux.NewRouter()

	// Rota para retornar uma mensagem de boas-vindas
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bem-vindo à minha API Go com Gorilla Mux!")
	})

	// Rota para listar todos os itens (exemplo de rota)
	r.HandleFunc("/items", ListItems).Methods("GET")

	// Rota para criar um novo item (exemplo de rota)
	r.HandleFunc("/items", CreateItem).Methods("POST")

	// Rota para obter detalhes de um item (exemplo de rota)
	r.HandleFunc("/items/{id}", GetItem).Methods("GET")

	// Rota para atualizar um item (exemplo de rota)
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")

	// Rota para excluir um item (exemplo de rota)
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	// Inicie o servidor HTTP
	http.Handle("/", r)
	fmt.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", nil)
}

// Função para listar todos os itens (simulação)
func ListItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Listando todos os itens")
}

// Função para criar um novo item (simulação)
func CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Criando um novo item")
}

// Função para obter detalhes de um item (simulação)
func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	fmt.Fprintf(w, "Detalhes do item %s\n", itemID)
}

// Função para atualizar um item (simulação)
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	fmt.Fprintf(w, "Atualizando o item %s\n", itemID)
}

// Função para excluir um item (simulação)
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	fmt.Fprintf(w, "Excluindo o item %s\n", itemID)
}
