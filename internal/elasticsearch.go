package internal

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// Logger contains function for logging.
type Logger interface {
	Send(key string, data interface{}) error
}

type client struct {
	es *elasticsearch.Client
}

// NewES to create new elasticsearch client.
func NewES(addresses []string, username string, password string) (Logger, error) {
	return NewESWithConfig(elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
	})
}

// NewESWithConfig to create new elasticsearch client with config.
func NewESWithConfig(cfg elasticsearch.Config) (Logger, error) {
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// Connection check.
	if err = isError(es.Info()); err != nil {
		return nil, err
	}
	return &client{es: es}, nil
}

// Send to send data to elasticsearch.
func (c *client) Send(key string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index:   key,
		Body:    strings.NewReader(string(d)),
		Refresh: "true",
	}
	return isError(req.Do(context.Background(), c.es))
}

func isError(res *esapi.Response, err error) error {
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}
