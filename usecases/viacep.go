package uscviacep

import (
	"busca_cep/domain"
	"encoding/json"
	"io"
	"net/http"
)

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
