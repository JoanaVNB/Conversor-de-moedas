package cron

import (
	"encoding/csv"
	"encoding/json"
	"sync"

	"exchange/presenter"
	"exchange/repository"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/robfig/cron"
	"gorm.io/gorm"
)

var dados []presenter.Presenter
var dadosMutex sync.Mutex 

func AutomatedRoutine(db *gorm.DB) {
	arquivo, err := os.OpenFile("dados.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer arquivo.Close()

	cron := cron.New()

	saveData := func() {
		client := http.Client{}
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/exchange/10/BRL/USD/4.50", nil)
		if err != nil {
				fmt.Println("Error when making HTTP request:", err)
				return
		}

		resp, err := client.Do(req)
		if err != nil {
				fmt.Println("Error capturing HTTP response:", err)
				return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
				fmt.Println("Error parsing JSON response:", err)
				return
		}

		var response map[string]interface{}
		err = json.Unmarshal(body, &response)
		if err != nil {
				fmt.Println(err)
				return
		}

		dadosMutex.Lock()
		defer dadosMutex.Unlock()

		valorConvertido := response["valorConvertido"].(float64)
		simboloMoeda := response["simboloMoeda"].(string)
		dados = append(dados, presenter.Presenter{ValorConvertido: valorConvertido, SimboloMoeda: simboloMoeda})

		writer := csv.NewWriter(arquivo)
		for _, dado := range dados {
				valueConverted := fmt.Sprintf("%.6f", dado.ValorConvertido)
				writer.Write([]string{valueConverted, dado.SimboloMoeda,})
		}
		writer.Flush()
		repository.Delete(db)
		fmt.Println("Dados salvos no arquivo e deletado da base de dados", time.Now().Format("02/01/2006 15:04:05"))

	}
 
	cron.AddFunc("1 * * * *", saveData)
	cron.Start()
	select {}
}



