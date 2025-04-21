package initialize

import (
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"log"
)

var ESClient *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response from Elasticsearch: %s", err)
	}
	defer res.Body.Close()

	fmt.Println("Elasticsearch connected successfully")
	ESClient = client
}
