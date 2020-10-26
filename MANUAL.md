# GO FRETE CORREIOS
API em Golang para consulta de valor de frete pelos correios

## Uso
Basta enviar uma requisição POST para a url ```http://localhost:3000/frete``` O retorno será um JSON com o conteúdo:

com as seguintes informações, exemplo:
* codigo_servido_desejado: 41106
* cep_origem: 11680000
* cep_destino: 82220000
* peso: 1
* altura: 15
* largura: 22
* comprimento: 32
* valor_produto: 0

O retorno será:
```
{
  "frete": {
    "codigo": "41106",
    "valor": "29,40",
    "prazoEntrega": "14",
    "valorSemAdicionais": "29,40",
    "valorNaoPropria": "0,00",
    "valorAvisoRecebimento": "0,00",
    "valorDeclarado": "0,00",
    "entregaDomicializar": "S",
    "entregaSabado": "N"
  }
}
```

## Código do serviço desejado

### Código comuns utilizados em loja virtual
* 40215 = SEDEX 10
* 40045 = SEDEX À COBRAR
* 40010 = SEDEX
* 41106 = PAC

### Todos os códigos disponíveis pelos nos Correrios
* 84050 = FAC Registrado
* 82031 = FAC Simples
* 20010 = Impresso Normal
* 14036 = Mala Direta Domiciliária
* 14010 = Mala Direta Básica e Especial
* 44105 = Malote
* 04510 = PAC
* 04693 = PAC Grandes Formatos
* 04707 = PAC Pagamento na Entrega
* 43010 = Reembolso Postal 
* 36200 = Remessa Econômica 
* 40355 = Remessa Expressa - CRLV/CRV/CNH 
* 40622 = Remessa Expressa Talão/Cartão 
* 75043 = Remessa Local com Comprovação de Entrega
* 04014 = SEDEX
* 40215 = SEDEX 10
* 40169 = SEDEX 12
* 04138 = SEDEX Contrato Grandes Formatos     
* 40290 = SEDEX HOJE
* 03662 = SEDEX HOJE COLABORAT A FATURAR 
* 03670 = SEDEX HOJE DESVIO COLAB A FATURAR 
* 04065 = SEDEX Pagamento na Entrega


## Configurar porta onde a aplicação irá executar
Para configurar a porta da aplicação ```:3003``` para outra porta como ```:8181```, 
basta abrir o arquivo **.env** e alterar a configuração: ```LOCALHOST_PORT=3003```

## Instalação da aplicação em sistema linux baseado em debian
Para colocar a aplicação para ser executada em produção em um servidor ubuntu server, 
basta executar o arquivo *install.sh* como root, exemplo: ```sudo ./install```

Feito isso a seguinte mensagem deve aparecer:
```
GO-FRETE - Serviço de consulta de endereço por FRETE

[√] Arquivos da aplicação copiado para /usr/local/gofrete/
[√] Permissões setadas
[√] Arquivo go-frete-correios.service copiado para /etc/systemd/system/go-frete-correios.service
[√] Atualizando lista de serviços do sistema operacional
[√] Executando serviço


lines 1--1...skipping...
● go-frete-correios.service
   Loaded: loaded (/etc/systemd/system/go-frete-correios.service; disabled; vendor preset: enabled)
   Active: active (running) since Mon 2020-10-26 08:48:25 -03; 40ms ago
 Main PID: 24075 (go-frete-correi)
    Tasks: 1 (limit: 4915)
   CGroup: /system.slice/go-frete-correios.service
           └─24075 /usr/local/gofrete/go-frete-correios

Fim da instalação, em caso de problema execute o comando: [ journalctl -u go-frete-correios -f ] para mais detalhes
```
Ao executar o comando *journalctl -u go-frete-correios -f* verá em qual endereço e porta a aplicação está acessível

```
Sep 07 17:48:12 tayron-ThinkPad-T420 systemd[1]: Started go-frete-correios.service.
Sep 07 17:48:12 tayron-ThinkPad-T420 go-frete-correios[5437]: Arquivo de configuração: /usr/local/gofrete/.env
Sep 07 17:48:12 tayron-ThinkPad-T420 go-frete-correios[5437]: Servidor executando no endereço: http://127.0.0.1:3003
```

Conforme o log a aplicação está sendo executada no endereço: ```http://127.0.0.1:3003```

## Desinstalação da aplicação em sistema linux baseado em debian
Para desinstalar basta executar o comando: ```sudo ./uninstall.sh```

Feito isso a seguinte mensagem deve aparecer:
```
Desinstalando GO-FRETE - Serviço de consulta de endereço por FRETE

[√] Parando serviço
[√] Removendo serviço antigo
[√] Atualizando lista de serviços do sistema operacional

Desinstalação concluída
```

## Exemplos de uso
Dentro da pasta *exemplos-uso* você encontrará exemplo de utilização da API nas linguagens JAVA, Javascript, PHP e Python.