package es

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

type EsHelper struct {
	username string
	password string
	hosts    []string
	sniff    bool

	requests []elastic.BulkableRequest
}

func Init(hosts []string) *EsHelper {
	return &EsHelper{
		hosts: hosts,
	}
}

func Helper() *EsHelper {
	return &EsHelper{}
}

func (helper *EsHelper) SetUsername(username string) {
	if username == "" {
		return
	}
	helper.username = username
}

func (helper *EsHelper) SetPassword(password string) {
	if password == "" {
		return
	}
	helper.password = password
}

func (helper *EsHelper) SetNiff() {
	helper.sniff = true
}

func (helper *EsHelper) GetClient() (*elastic.Client, error) {
	if len(helper.hosts) == 0 {
		return nil, errors.New("es hosts is empty")
	}

	options := []elastic.ClientOptionFunc{
		elastic.SetURL(helper.hosts...),
	}

	if helper.sniff {
		options = append(options,
			elastic.SetSniff(true))
	}

	if helper.username != "" && helper.password != "" {
		options = append(options,
			elastic.SetBasicAuth(helper.username, helper.password))
	}

	client, err := elastic.NewClient(options...)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (helper *EsHelper) IndexRequest(esindex string, estype string, data interface{}) *elastic.BulkIndexRequest {
	return elastic.NewBulkIndexRequest().Index(esindex).Type(estype).Doc(data)
}

func (helper *EsHelper) DeleteRequest(esindex string, estype string, id string) *elastic.BulkDeleteRequest {
	return elastic.NewBulkDeleteRequest().Index(esindex).Type(estype).Id(id)
}

func (helper *EsHelper) AddRequest(req elastic.BulkableRequest) {
	if len(helper.requests) == 0 {
		helper.requests = make([]elastic.BulkableRequest, 0)
	}
	helper.requests = append(helper.requests, req)
}

func (helper *EsHelper) GetRequestsAndClear() []elastic.BulkableRequest {
	r := helper.requests
	helper.clear()
	return r
}

func (helper *EsHelper) clear() {
	helper.requests = helper.requests[0:0]
}

func (helper *EsHelper) Bulk(client *elastic.Client, requests ...elastic.BulkableRequest) (*elastic.BulkResponse, error) {
	service := client.Bulk().Add(requests...)
	return service.Do(context.TODO())
}
