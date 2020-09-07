package config

import (
	"net/http"

	"github.com/tayron/go-frete-correios/application/controller"

	"github.com/gorilla/mux"
)

func LoadHouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Home)
	router.HandleFunc("/frete", controller.GetFrete).Methods("POST")
	http.Handle("/", router)
}
