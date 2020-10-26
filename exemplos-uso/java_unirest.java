Unirest.setTimeouts(0, 0);
HttpResponse<String> response = Unirest.post("http://127.0.0.1:3003/frete")
  .multiPartContent()  .field("codigo_servido_desejado", "41106")
  .field("cep_origem", "11680000")
  .field("cep_destino", "82220000")
  .field("peso", "1")
  .field("altura", "15")
  .field("largura", "22")
  .field("comprimento", "32")
  .field("valor_produto", "0")
  .asString();
