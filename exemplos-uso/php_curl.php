<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => "http://127.0.0.1:3003/frete",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => "",
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => "POST",
  CURLOPT_POSTFIELDS => array('codigo_servido_desejado' => '41106','cep_origem' => '11680000','cep_destino' => '82220000','peso' => '1','altura' => '15','largura' => '22','comprimento' => '32','valor_produto' => '0'),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;