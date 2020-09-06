package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rogpeppe/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"
	"github.com/tayron/go-cep/application/entity"
)

type Frete struct {
	CServico struct {
		Codigo                string `xml:"Codigo"`
		Valor                 string `xml:"Valor"`
		PrazoEntrega          string `xml:"PrazoEntrega"`
		ValorSemAdicionais    string `xml:"ValorSemAdicionais"`
		ValorMaoPropria       string `xml:"ValorMaoPropria"`
		ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento"`
		ValorValorDeclarado   string `xml:"ValorValorDeclarado"`
		EntregaDomiciliar     string `xml:"EntregaDomiciliar"`
		EntregaSabado         string `xml:"EntregaSabado"`
	} `xml:"cServico" json:"frete"`
}

const urlCorrerios = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?nCdEmpresa=&sDsSenha=&sCepOrigem=%s&sCepDestino=%s&nVlPeso=%s&nCdFormato=1&nVlComprimento=%s&nVlAltura=%s&nVlLargura=%s&sCdMaoPropria=n&nVlValorDeclarado=%s&sCdAvisoRecebimento=n&nCdServico=%s&nVlDiametro=0&StrRetorno=xml"

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

func CalcularFrete(parametros entity.ParametroCorrerios) (Frete, error) {

	var frete Frete
	urlAPICorrerios := montarUrlApiCorrerios(parametros)

	req, err := http.Get(urlAPICorrerios)
	if err != nil {
		return frete, err
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return frete, err
	}

	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader

	errDecode := decoder.Decode(&frete)

	if errDecode != nil {
		return frete, errDecode
	}

	return frete, nil
}
