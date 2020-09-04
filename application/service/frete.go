package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tayron/go-cep/application/entity"
)

type Frete struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

const urlCorrerios = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?nCdEmpresa=&sDsSenha=&sCepOrigem=%s&sCepDestino=%s&nVlPeso=%s&nCdFormato=1&nVlComprimento=%s&nVlAltura=%s&nVlLargura=%s&sCdMaoPropria=n&nVlValorDeclarado=%s&sCdAvisoRecebimento=n&nCdServico=%s&nVlDiametro=0&StrRetorno=json"

func montarUrlApiCorrerios(parametros entity.ParametroCorrerios) string {
	return fmt.Sprintf(urlCorrerios,
		parametros.CepOrigem,
		parametros.CepDestino,
		parametros.Peso,
		parametros.Comprimento,
		parametros.Altura,
		parametros.Largura,
		parametros.ValorProduto,
		parametros.CodigoServicoDesejado,
	)
}

func CalcularFrete(parametros entity.ParametroCorrerios) (string, error) {

	urlAPICorrerios := montarUrlApiCorrerios(parametros)

	fmt.Println(urlAPICorrerios)

	req, err := http.Get(urlAPICorrerios)
	if err != nil {
		return "", err
	}

	var frete Frete
	err = json.NewDecoder(req.Body).Decode(&frete)
	if err != nil {
		return "", err
	}
	res, err := json.Marshal(frete)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
