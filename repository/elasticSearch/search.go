package elasticSearch

import (
	"context"
	"encoding/json"
	"exchange/presenter"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)


func SeachValue(ctx context.Context, value string) {

	client := ctx.Value(ClientKey).(*elasticsearch.Client)

	response, err := client.Get("datas", value)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer func() {
		if response.Body != nil {
			_ = response.Body.Close()
		}
	}()

	if response.IsError(){
		log.Fatalf("Document with this value %s was not found", value)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
			log.Fatalf("Error reading response: %s", err)
	}
	fmt.Printf("Response: %s\n", responseBytes)

	type GetData struct {
		Index string `json:"_index"`
		Type string `json:"_type"`
		ID   string `json:"_id"`
		Version int  `json:"_version"`
		Seq int `json:"_seq_no"`
		Term int `json:"_primary_term"`
		Found bool `json:"found"`
		Source *presenter.Presenter	 `json:"_source"`
		CreatedAt time.Time `json:"_created_at"`
		UpdatedAt time.Time `json:"_updated_at"`
}

var getData GetData
	err = json.Unmarshal(responseBytes, &getData)
	if err != nil {
    log.Fatalf("Error decoding response: %s", err)
}

/* 	err = json.NewDecoder(response.Body).Decode(&getUser)
	//fmt.Printf("GetResponse: %s", getResponse)
	if err != nil {
		log.Fatalf("Error to decode: %s", err)
	} */
	

	fmt.Printf("✅ Data with the value %s was found\n", value)
}