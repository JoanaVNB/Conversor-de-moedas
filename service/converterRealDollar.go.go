package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RateUSD() (ptax float64){
	url := "https://economia.awesomeapi.com.br/last/USD-BRL"

	res, err := http.Get(url)
	if err != nil {
			fmt.Println("Erro ao fazer a solicitação HTTP:", err)
	}
	defer res.Body.Close()
		
	// Analisar a resposta JSON da API
	type Response struct {
			Rates struct {
					Bid string `json:"bid"`
					Ask	string `json:"ask"`
			} `json:"USDBRL"`
	}

	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
			fmt.Println("Erro ao analisar a resposta JSON:", err)
			return 0
	}

	bid, err := strconv.ParseFloat(response.Rates.Bid, 64)
	if err != nil {
			fmt.Println("Erro ao converter bid para float:", err)
			return 0
	}

	ask, err := strconv.ParseFloat(response.Rates.Ask, 64)
	if err != nil {
			fmt.Println("Erro ao converter ask para float:", err)
			return 0
	}

	ptax = (bid + ask) / 2
	return ptax
}

func BRLconvertUSD(real float64, rateUSD float64) (dolar float64){
	dolar = real / rateUSD
	return dolar
}

func USDconvertBRL(dolar float64, rateUSD float64) (real float64){
	real = dolar/ rateUSD
	return real
}
