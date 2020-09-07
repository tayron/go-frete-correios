package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tayron/go-cep/application/entity"
	"github.com/tayron/go-cep/application/service"
)

func Test_CalcularFrete(t *testing.T) {

	t.Run("invalid frete", func(t *testing.T) {
		var parametros entity.ParametroCorrerios = entity.ParametroCorrerios{
			CodigoServicoDesejado: "4110699999",
			CepOrigem:             "11680000",
			CepDestino:            "82220000",
			Peso:                  "1",
			Altura:                "15",
			Largura:               "22",
			Comprimento:           "32",
			ValorProduto:          "0",
		}

		frete, err := service.CalcularFrete(parametros)

		assert.Nil(t, err)
		assert.Equal(t, "4110699999", frete.CServico.Codigo)
		assert.Equal(t, "0", frete.CServico.PrazoEntrega)
		assert.Equal(t, "0,00", frete.CServico.Valor)
	})

	t.Run("valid frete", func(t *testing.T) {
		var parametros entity.ParametroCorrerios = entity.ParametroCorrerios{
			CodigoServicoDesejado: "41106",
			CepOrigem:             "11680000",
			CepDestino:            "82220000",
			Peso:                  "1",
			Altura:                "15",
			Largura:               "22",
			Comprimento:           "32",
			ValorProduto:          "0",
		}

		frete, err := service.CalcularFrete(parametros)

		assert.Nil(t, err)
		assert.Equal(t, "41106", frete.CServico.Codigo)
	})
}
