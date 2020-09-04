package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tayron/go-cep/application/entity"

	"github.com/gorilla/mux"
	"github.com/tayron/go-cep/application/service"
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
	vars := mux.Vars(req)
	rw.Header().Set("Content-Type", "application/json")

	parametrosCorrerios := entity.ParametroCorrerios{
		CodigoServicoDesejado: vars["codigo_servido_desejado"],
		CepOrigem:             vars["cep_origem"],
		CepDestino:            vars["cep_destino"],
		Peso:                  vars["peso"],
		Altura:                vars["altura"],
		Largura:               vars["largura"],
		Comprimento:           vars["comprimento"],
		ValorProduto:          vars["valor_produto"],
	}

	frete, err := service.CalcularFrete(parametrosCorrerios)
	if err != nil {
		respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
		return
	}
	_, err = rw.Write([]byte(frete))
	if err != nil {
		respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
		return
	}
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
