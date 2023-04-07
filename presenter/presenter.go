package presenter

import "exchange/domain"

type Presenter struct{
	ValorConvertido	float64 `json:"valorConvertido"`
	SimboloMoeda	string	`json:"simboloMoeda"`
}

func PresenterFX(e domain.Exchange) *Presenter{
	return &Presenter{
		ValorConvertido:	e.ValueConverted,
		SimboloMoeda:		e.SymbolTo,
	}
}