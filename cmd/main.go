package main

import (
	"busca_cep/domain"
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!!!"))
}

func BuscaCep(cep string) (address *domain.ViaCEP, err error) {
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return
	}
	defer res.Body.Close()

	bodyResult, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyResult, &address)
	if err != nil {
		return
	}

	return
}
