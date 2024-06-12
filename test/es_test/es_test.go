package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func TestCreateIndex(t *testing.T) {
	// Set up the Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{"https://8af9afd9e4bf4d88b97b14488467361d.us-central1.gcp.cloud.es.io"},
		APIKey:    "SEZ3cUI1QUJaclpXZ01wZGhPckE6UlRkcjZXeENRQjJXaEhISnF2eTBZQQ==",
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
func TestCreateQuestionIndex(t *testing.T) {
	// Set up the Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{"https://8af9afd9e4bf4d88b97b14488467361d.us-central1.gcp.cloud.es.io"},
		APIKey:    "SEZ3cUI1QUJaclpXZ01wZGhPckE6UlRkcjZXeENRQjJXaEhISnF2eTBZQQ==",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		t.Fatalf("Error creating the client: %s", err)
	}

	// 删除现有的索引（如果有的话）
	res, err := es.Indices.Delete([]string{"questions"}, es.Indices.Delete.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("Cannot delete index: %s", err)
	}
	defer res.Body.Close()

	// 索引映射
	mapping := `{
		"mappings": {
			"properties": {
				"question_id": {
					"type": "integer"
				},
				"question_type": {
					"type": "integer"
				},
				"question_difficulty": {
					"type": "integer"
				},
				"question_content": {
					"type": "text"
				},
				"question_answer": {
					"type": "text"
				},
				"question_grade": {
					"type": "integer"
				},
				"options": {
					"type": "object",
					"enabled": true
				}
			}
		}
	}`

	// 创建索引
	req := esapi.IndicesCreateRequest{
		Index: "questions",
		Body:  bytes.NewReader([]byte(mapping)),
	}

	res, err = req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error creating index: %s", res.String())
	} else {
		fmt.Println("Index created successfully")
	}
}
