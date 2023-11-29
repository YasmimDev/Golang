package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey Gopher"))
}

func main() {
	// HTTP É UM SERVIDOR DE COMUNICAÇÃO - BASE DE COMUNICAÇÃO WEB
	// CLIENTE (faz requisição) - SERVIDOR (processa requisição e envia resposta)
	// Request - Response
	// Rotas
	// URI - identificador de recursos
	// metodo - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
