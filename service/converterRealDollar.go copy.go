package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RateEUR() (ptax float64){
	url := "http://economia.awesomeapi.com.br/json/last/EUR-BRL"

	res, err := http.Get(url)
    if err != nil {
        fmt.Println("Erro ao fazer a solicitação HTTP:", err)
        return 0
    }
    defer res.Body.Close()

    type Response struct {
			Rates struct {
					Bid string `json:"bid"`
					Ask string `json:"ask"`
			} `json:"EURBRL"`
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








func BRLconvertEUR(real float64, rateEUR float64) (euro float64){
	euro = real / rateEUR
	return euro
}

func EURconvertBRL(euro float64, rateEUR float64) (real float64){
	real = euro/ rateEUR
	return real
}
