# Conversor-de-moedas

## Descrição do projeto
API Rest para conversão de moedas, com restrição para as seguintes:
* De Real para Dólar; De Dólar para Real;
* De Real para Euro; De Euro para Real;
* De BTC para Dolar; De BTC para Real;

Utilize `docker-compose up -d`. Após os containers estarem de pé, execute a aplicação com `go run main.go` . 

## Versões

* v1 - a conversão é realizada apenas com taxa definida
* v2 - a conversão também é realizada com taxa real do mercado
* v3- implementação do docker
* v4 - sem docker, e domain sem utilidade(tá certo isso?)

## Requisitos:
**Banco de dados**: salvar os dados no banco de dados MySQL e criar uma rotina para salvar o dados para consultas futuras:
Para isto foi criado o diretório "repository", em que foi utilizado o framework ORM de Golang, Gorm, e onde contém os códigos para:
* armazenar a entidade Exchange;
* retorna valor cotado, a partir de uma taxa de conversão real ou definida pelo usuário;
* poder fazer a busca por ID;
* por moeda a ser convertida;
* por todos os dados cadastrados; e
* para detelar.

Para criação de **rotina de armazenamento** de dados, foi utilizado o cron, em que os dados são salvos, com intervalo de 1 minuto, a partir do retorno da url: http://localhost:8000/exchange/10/BRL/USD/4.50 no arquivo "dados.csv" e são apagados do banco de dados logo após o armazenamento no arquivo, para evitar dados não-realistas. Tentei buscar também com ElasticSearch, mas não consegui arrumar a tempo.

O armazenamento no arquivo foi feito a partir do que era esperado como mensagem de retorno:

```
{
  "valorConvertido": 45,
  "simboloMoeda": "$"
}
```

Enquanto no banco de dados, dei preferência para salvar mais informações, de acordo com a entidade Exchange.

## Endpoints
A URL da requisição deve seguir o seguinte formato:
* http://localhost:8000/exchange/{amount}/{from}/{to}/{rate}
* http://localhost:8000/exchange/10/BRL/USD/4.50

Foram criados os seguintes endpoints:
* GET -> Faz a conversão, salva dados e mostra retorno a partir de uma taxa de conversão definida:
  "/exchange/:amount/:from/:to/:rate"
  http://localhost:8000/exchange/10/BRL/USD/4.50
* GET -> Faz a conversão, salva dados e mostra retorno a partir de uma taxa de conversão real:
  "/exchange/:amount/:from/:to"
  http://localhost:8000/exchange/10/BRL/USD
*	GET -> Busca por ID:
  "/exchange/id/:id"
  http://localhost:8000/exchange/id/1
*	GET -> Busca por todos os cadastros:
  "/exchange"
  http://localhost:8000/exchange
*	GET -> Busca por cadastro que foram convertidos a partir de uma específica moeda:
    "/exchange/from/:from"
    http://localhost:8000/exchange/from/BRL
*	DELETE -> Deleta cadastro:
     "/exchange/id/:id"
     http://localhost:8000/exchange/id/1

## Testes
Para testar, use o comando `go test` no diretório que deseja rodar os testes.
Os testes foram realizados com o framework GoMock.

