package main

import (
	"busca_cep/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.BuscaCep)
	http.ListenAndServe(":8080", mux)
}
