package main

import (
	uscviacep "busca_cep/usecases"
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCep)
	http.ListenAndServe(":8080", nil)
}

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := uscviacep.BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}
