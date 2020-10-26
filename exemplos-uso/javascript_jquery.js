var form = new FormData();
form.append("codigo_servido_desejado", "41106");
form.append("cep_origem", "11680000");
form.append("cep_destino", "82220000");
form.append("peso", "1");
form.append("altura", "15");
form.append("largura", "22");
form.append("comprimento", "32");
form.append("valor_produto", "0");

var settings = {
  "url": "http://127.0.0.1:3003/frete",
  "method": "POST",
  "timeout": 0,
  "processData": false,
  "mimeType": "multipart/form-data",
  "contentType": false,
  "data": form
};

$.ajax(settings).done(function (response) {
  console.log(response);
});