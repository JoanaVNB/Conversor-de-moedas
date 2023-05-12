package service

import (
	"exchange/domain"
	"time"
)

type Exchange struct {
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

//não foi utilizado
func NewExchange(ex domain.Exchange) *Exchange{
	return &Exchange{
		ID: ex.ID,
		From: ex.From,
		To: ex.To,
		SymbolTo: ex.SymbolTo,
		SymbolFrom: ex.SymbolFrom,
		ValueFrom: ex.ValueFrom,
		Rate: ex.Rate, 
		ValueConverted: ex.ValueConverted, 
		Converted: ex.Converted, 
		CreatedAt: ex.CreatedAt,
	}
}


//Dessa forma, a struct Exchange no pacote service será um alias da struct
// Exchange no pacote domain. Todas as definições e campos da struct 
// Exchange no pacote service serão idênticos aos da struct Exchange no 
// pacote domain. Qualquer alteração na struct Exchange em domain será 
// refletida automaticamente na struct Exchange em service.