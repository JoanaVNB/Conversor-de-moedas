package main

import (
	"exchange/service"
	"fmt"
)

func main(){
	rateUSD := service.RateUSD()
	dolar := service.BRLconvertUSD(100, rateUSD)
	fmt.Println(dolar)

	rateEUR := service.RateEUR()
	euro := service.BRLconvertEUR(100, rateEUR)
	fmt.Println(euro)

	rateBTC := service.RateBTC()
	bitcoin := service.BRLconvertBTC(100, rateBTC)
	fmt.Println(bitcoin)

}