package lto

import (
	"github.com/go-resty/resty/v2"
)

func NewAPI(config *Config) (*API, error) {
	client := resty.New()
	client.SetHostURL(config.NodeAddress)

	return &API{
		client: client,
		config: config,
	}, nil
}

type API struct {
	client *resty.Client
	config *Config
}
