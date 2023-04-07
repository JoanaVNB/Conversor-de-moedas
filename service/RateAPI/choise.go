package rate

//import "exchange/domain"

//Função para escolher para qual função indicar
func ChoiseFX(amount float64, from string, to string) (result float64, rate float64) {
	if from == "BRL" && to == "USD" {
		result, rate = BRLconvertUSD(amount)
	}
	if from == "USD" && to == "BRL" {
		result, rate  = USDconvertBRL(amount)
	}
	if from == "BRL" && to == "EUR" {
		result, rate  = BRLconvertEUR(amount)
	}
	if from == "EUR" && to == "BRL" {
		result, rate  = EURconvertBRL(amount)
	}
	if from == "BTC" && to == "USD" {
		result, rate  = BTCconvertUSD(amount)
	}
	if from == "BTC" && to == "BRL" {
		result, rate  = BTCconvertBRL(amount)
	}
	return result, rate 
}
