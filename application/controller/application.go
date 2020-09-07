package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tayron/go-frete-correios/application/entity"

	"github.com/tayron/go-frete-correios/application/service"
)

const errorMessage = "Erro ao consultar frete"

func Home(rw http.ResponseWriter, req *http.Request) {
	_, err := rw.Write([]byte([]byte("ping")))
	if err != nil {
		respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
		return
	}
}

func GetFrete(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	parametrosCorrerios := entity.ParametroCorrerios{
		CodigoServicoDesejado: req.FormValue("codigo_servido_desejado"),
		CepOrigem:             req.FormValue("cep_origem"),
		CepDestino:            req.FormValue("cep_destino"),
		Peso:                  req.FormValue("peso"),
		Altura:                req.FormValue("altura"),
		Largura:               req.FormValue("largura"),
		Comprimento:           req.FormValue("comprimento"),
		ValorProduto:          req.FormValue("valor_produto"),
	}

	frete, err := service.CalcularFrete(parametrosCorrerios)

	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(frete)
}

//RespondWithError return a http error
func respondWithError(w http.ResponseWriter, code int, e string, message string) {
	respondWithJSON(w, code, map[string]string{"code": strconv.Itoa(code), "error": e, "message": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
