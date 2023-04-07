package rate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RateBTC() (ptax float64){
    url := "http://economia.awesomeapi.com.br/json/last/BTC-BRL"

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
        } `json:"BTCBRL"`
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

func RateBTC_USD() (ptax float64){
    url := "http://economia.awesomeapi.com.br/json/last/BTC-USD"

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
        } `json:"BTCUSD"`
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

func BTCconvertUSD(amountBitcoin float64) (dolar float64, rate float64){
    rate = RateBTC_USD()
    dolar = amountBitcoin * rate
    return dolar, rate 
}

func BTCconvertBRL(amountBitcoin float64) (real float64, rate float64){
    rate = RateBTC()
	real = amountBitcoin * rate
	return real, rate
}
