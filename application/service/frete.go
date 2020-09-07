package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rogpeppe/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"
	"github.com/tayron/go-frete-correios/application/entity"
)

type Frete struct {
	CServico struct {
		Codigo                string `xml:"Codigo" json:"codigo"`
		Valor                 string `xml:"Valor" json:"valor"`
		PrazoEntrega          string `xml:"PrazoEntrega" json:"prazoEntrega"`
		ValorSemAdicionais    string `xml:"ValorSemAdicionais" json:"valorSemAdicionais"`
		ValorMaoPropria       string `xml:"ValorMaoPropria" json:"valorNaoPropria"`
		ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento" json:"valorAvisoRecebimento"`
		ValorDeclarado        string `xml:"ValorValorDeclarado" json:"valorDeclarado"`
		EntregaDomiciliar     string `xml:"EntregaDomiciliar" json:"entregaDomicializar"`
		EntregaSabado         string `xml:"EntregaSabado" json:"entregaSabado"`
	} `xml:"cServico" json:"frete"`
}

const urlCorreios = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?nCdEmpresa=&sDsSenha=&sCepOrigem=%s&sCepDestino=%s&nVlPeso=%s&nCdFormato=1&nVlComprimento=%s&nVlAltura=%s&nVlLargura=%s&sCdMaoPropria=n&nVlValorDeclarado=%s&sCdAvisoRecebimento=n&nCdServico=%s&nVlDiametro=0&StrRetorno=xml"

func montarUrlApiCorreios(parametros entity.ParametroCorreios) string {
	return fmt.Sprintf(urlCorreios,
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

func CalcularFrete(parametros entity.ParametroCorreios) (Frete, error) {

	var frete Frete
	urlAPICorreios := montarUrlApiCorreios(parametros)

	req, err := http.Get(urlAPICorreios)
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
