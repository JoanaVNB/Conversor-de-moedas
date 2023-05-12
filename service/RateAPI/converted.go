package rate

import (
	"exchange/service"
	"fmt"
)

func GetConverted(amount float64, from string, to string, rate float64, result float64) service.Exchange{
	symbolFrom, _ := service.GetCurrencySymbol(from)
	symbolTo, _ := service. GetCurrencySymbol(to)
	converted := fmt.Sprintf("%s%.2f->%s%.2f", symbolFrom, amount, symbolTo, result)

	ex := service.Exchange{
		From:           from,
		To:             to,
		SymbolTo:       symbolTo,
		SymbolFrom:     symbolFrom,
		ValueFrom:      amount,
		Rate: 					rate,		
		ValueConverted: 	result,
		Converted:      converted,
	}
	return ex
}



