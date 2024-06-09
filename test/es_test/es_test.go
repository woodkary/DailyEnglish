package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func TestCreateIndex(t *testing.T) {
	// Set up the Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{"https://b8fde32e62044f12b769b107e7e2346f.us-central1.gcp.cloud.es.io"},
		APIKey:    "RzQ0VC1JOEJMQ0gwOXRlMFloZkQ6M2dpNUFMRF9SeE9wMkxhNjAxUjF5dw==",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		t.Fatalf("Error creating the client: %s", err)
	}

	// Define the index mapping
	mapping := map[string]interface{}{
		"settings": map[string]interface{}{
			"number_of_shards":   1, //分片数
			"number_of_replicas": 1, //副本数
		},
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"word_id": map[string]interface{}{
					"type": "integer",
				},
				"spelling": map[string]interface{}{
					"type": "text",
				},
				"pronunciation": map[string]interface{}{
					"type": "text",
				},
				"meanings": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"verb": map[string]interface{}{
							"type": "keyword",
						},
						"noun": map[string]interface{}{
							"type": "keyword",
						},
						"pronoun": map[string]interface{}{
							"type": "keyword",
						},
						"adjective": map[string]interface{}{
							"type": "keyword",
						},
						"adverb": map[string]interface{}{
							"type": "keyword",
						},
						"preposition": map[string]interface{}{
							"type": "keyword",
						},
						"conjunction": map[string]interface{}{
							"type": "keyword",
						},
						"interjection": map[string]interface{}{
							"type": "keyword",
						},
					},
				},
			},
		},
	}

	// Create the index with the specified mapping
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(mapping); err != nil {
		t.Fatalf("Error encoding mapping: %s", err)
	}

	req := esapi.IndicesCreateRequest{
		Index:         "dailyenglish",
		Body:          &buf,
		MasterTimeout: time.Second * 60,
		Timeout:       time.Second * 30,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		t.Fatalf("Error creating the index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		t.Fatalf("Error response from Elasticsearch: %s", res.String())
	} else {
		fmt.Println("Index created successfully")
	}
}
