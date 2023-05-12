package service

import (
	"fmt"
)

func GetCurrencySymbol(code string) (converted string, err error){
	currencySymbols := map[string]string{
		"BRL": "R$",
		"USD": "$",
		"EUR": "€",
		"BTC": "₿",
	}
	symbol, ok := currencySymbols[code]
	if !ok {
			return "", fmt.Errorf("currency %s not found", code)
	}
	return symbol, nil
}

func GetConverted(amount float64, from string, to string, rate float64) Exchange{
	symbolFrom, _ := GetCurrencySymbol(from)
	symbolTo, _ := GetCurrencySymbol(to)
	valueConverted := amount / rate
	converted := fmt.Sprintf("%s%.2f->%s%.2f", symbolFrom, amount, symbolTo, valueConverted)

	ex := Exchange{
		From:           from,
		To:             to,
		SymbolTo:       symbolTo,
		SymbolFrom:     symbolFrom,
		ValueFrom:      amount,
		Rate: 					rate,		
		ValueConverted: valueConverted,
		Converted:      converted,
	}
	return ex
}



