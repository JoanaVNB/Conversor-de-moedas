package elasticSearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"context"
	"log"
	"fmt"
)


type ElasticSeach struct{
	Client *elasticsearch.Client
}

type contextKey struct {
	Key int
}

var DataKey contextKey = contextKey{Key: 1}
var ClientKey contextKey = contextKey{Key: 2}

func NewElasticSearch (Client *elasticsearch.Client,) *ElasticSeach{
	return &ElasticSeach{
		Client:   Client,
	}
}

func ConnectionWithElasticSearch(ctx context.Context) context.Context {
	newClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},})
	if err != nil{
		log.Fatalf("Error creating the client: %s", err)
	}
	fmt.Printf("✅ Client connected\n")
	return context.WithValue(ctx, ClientKey, newClient)
}
