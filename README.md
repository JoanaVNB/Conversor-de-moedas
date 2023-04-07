# Conversor-de-moedas

## Descrição do projeto
API Rest para conversão de moedas, com restrição para as seguintes:
* De Real para Dólar; De Dólar para Real;
* De Real para Euro; De Euro para Real;
* De BTC para Dolar; De BTC para Real;

## Requisitos:
**Banco de dados**: salvar os dados no banco de dados MySQL e criar uma rotina para salvar o dados para consultas futuras:
Para isto foi criado o diretório "repository", em que foi utilizado o framework ORM de Golang, Gorm, e onde contém os códigos para:
* armazenar a entidade Exchange;
* poder fazer a busca por ID;
* por moeda a ser convertida;
* por todos os dados cadastrados; e
* para detelar.

Para criação de **rotina de armazenamento** de dados, foi utilizado o cron, em que os dados são salvos a partir do retorno da url: http://localhost:8000/exchange/10/BRL/USD/4.50 no arquivo "dados.csv" e são apagados do banco de dados logo após o armazenamento no arquivo, para evitar dados não-realistas.

O armazenamento foi feito a partir do que era esperado como mensagem de retorno:

```
{
  "valorConvertido": 45,
  "simboloMoeda": "$"
}
```

Enquanto no banco de dados, dei preferência para salvar mais informações, de acordo com a entidade Exchange:

```
//ex: converter real para dólar
type Exchange struct{
	ID int	`json:"id" gorm:"primary key"`
	From string `json:"from" gorm:"varchar(3); not null"`//BRL
	To string `json:"to" gorm:"varchar(3); not null"`//USD
	SymbolTo string `json:"symbolTo" gorm:"varchar(3); not null"`//R$
	SymbolFrom string `json:"symboFrom" gorm:"varchar(3); not null"`//$
	ValueFrom float64 `json:"valueFrom" gorm:"not null"`//100
	Rate float64 `json:"rate" gorm:"not null"`//5
	ValueConverted	float64 `json:"valueConverted" gorm:"not null"`//20
	Converted	string `json:"converted" gorm:"not null"` //R$100 -> $20
	CreatedAt time.Time  `json:"created_at"`
}
```

## Endpoints
A URL da requisição deve seguir o seguinte formato:
* http://localhost:8000/exchange/{amount}/{from}/{to}/{rate}
* http://localhost:8000/exchange/10/BRL/USD/4.50

Foram criados os seguintes endpoints:
* GET -> Fazer a conversão, salvar dados, mostrar retorno:
  "/exchange/:amount/:from/:to/:rate"
  http://localhost:8000/exchange/10/BRL/USD/4.50
*	GET -> Busca por ID:
  "/exchange/id/:id"
  http://localhost:8000/exchange/id/1
*	GET -> Busca por todos os cadastros:
  "/exchange"
  http://localhost:8000/exchange
*	GET -> Busca por cadastro que foram convertidos a partir de uma específica moeda:
    "/exchange/from/:from"
    http://localhost:8000/exchange/from/BRL
*	DELETE -> Deletar cadastro:
     "/exchange/id/:id"
     http://localhost:8000/exchange/id/1

## Testes
Para testar, use o comando `go test` no diretório que deseja rodar os testes.
Os testes foram realizados com o framework GoMock.

