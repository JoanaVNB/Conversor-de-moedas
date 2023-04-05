# Conversor-de-moedas

## Descrição do projeto
API Rest para conversão de moedas.

## Objetivo:
Conversão das seguintes moedas:
* De Real para Dólar;
* De Dólar para Real;

* De Real para Euro;
* De Euro para Real;

* De BTC para Dolar;
* De BTC para Real;

## Requisitos:
* Salvar os dados no banco de dados MySQL;
* Criar uma rotina para salvar o dados para consultas futuras;

A URL da requisição deve seguir o seguinte formato:
* http://localhost:8000/exchange/{amount}/{from}/{to}/{rate}
* http://localhost:8000/exchange/10/BRL/USD/4.50

A resposta deve seguir o seguinte formato:
{
  "valorConvertido": 45,
  "simboloMoeda": "$"
}
