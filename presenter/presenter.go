package presenter

import "exchange/service"

type Presenter struct{
	ValorConvertido	float64 `json:"valorConvertido"`
	SimboloMoeda	string	`json:"simboloMoeda"`
}

func PresenterFX(e service.Exchange) *Presenter{
	return &Presenter{
		ValorConvertido:	e.ValueConverted,
		SimboloMoeda:		e.SymbolTo,
	}
}