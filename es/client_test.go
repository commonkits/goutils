package es

import (
	"context"
	"log"
	"testing"
)

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Test_Index(t *testing.T) {
	option := Init([]string{"http://192.168.2.194:9200"})
	option.SetUsername("elastic")
	option.SetPassword("Timevale#12345#")
	client, err := option.GetClient()
	if err != nil {
		return
	}

	resp, err := client.Index().Index("test").Type("type").BodyJson(&Test{Name: "safas", Age: 18}).Do(context.Background())
	if err != nil {
		log.Printf("index err: %v", err)
	}

	log.Printf("%v", resp)
}

func Test_bulk(t *testing.T) {
	esHelper := Init([]string{"http://192.168.2.194:9200"})
	esHelper.SetUsername("elastic")
	esHelper.SetPassword("Timevale#12345#")
	client, err := esHelper.GetClient()
	if err != nil {
		return
	}

	resp, err := esHelper.Bulk(client, esHelper.IndexRequest("test", "type", &Test{Name: "safas", Age: 18}))
	if err != nil {
		log.Printf("Bulk err: %v", err)
	}

	resp, err = esHelper.Bulk(client, esHelper.DeleteRequest("test", "type", "AWtvnUklqHlxAXPZousZ"))

	log.Printf("%v", resp)
}
