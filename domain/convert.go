package domain

import (
	"time"
)

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

