package lto

import (
	"context"
	"fmt"

	cryptorand "crypto/rand"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/unchainio/pkg/xapi"
)

var rand = cryptorand.Reader

func NewAPI(config *Config) (*API, error) {
	client, err := xapi.NewClient(config.NodeAddress)
	if err != nil {
		return nil, err
	}

	return &API{
		client: client,
		config: config,
	}, nil
}

type API struct {
	client *xapi.Client
	config *Config
}

type balanceResponse struct {
	Address       string `json:"address"`
	Confirmations int64  `json:"confirmations"`
	Balance       int64  `json:"balance"`
}

type BalanceResponse struct {
	Address       []byte `json:"address"`
	Confirmations int64  `json:"confirmations"`
	Balance       int64  `json:"balance"`
}

func (api *API) Balance(address []byte) (*BalanceResponse, error) {
	addressString := crypto.Base58Encode(address)
	res := new(balanceResponse)

	path := fmt.Sprintf("/addresses/balance/%s", addressString)
	_, _, _, err := api.client.Get(context.Background(), path, nil, res)
	if err != nil {
		return nil, err
	}

	return &BalanceResponse{
		Address:       crypto.Base58Decode(res.Address),
		Confirmations: res.Confirmations,
		Balance:       res.Balance,
	}, nil
}

func (api *API) BalanceWithConfirmations(address []byte, confirmations int) (*BalanceResponse, error) {
	addressString := crypto.Base58Encode(address)
	res := new(balanceResponse)

	path := fmt.Sprintf("/addresses/balance/%s/%d", addressString, confirmations)
	_, _, _, err := api.client.Get(context.Background(), path, nil, res)
	if err != nil {
		return nil, err
	}

	return &BalanceResponse{
		Address:       crypto.Base58Decode(res.Address),
		Confirmations: res.Confirmations,
		Balance:       res.Balance,
	}, nil
}

type balanceDetailsResponse struct {
	Address    string `json:"address"`
	Regular    int64  `json:"regular"`
	Generating int64  `json:"generating"`
	Available  int64  `json:"available"`
	Effective  int64  `json:"effective"`
}

type BalanceDetailsResponse struct {
	Address    []byte `json:"address"`
	Regular    int64  `json:"regular"`
	Generating int64  `json:"generating"`
	Available  int64  `json:"available"`
	Effective  int64  `json:"effective"`
}

func (api *API) BalanceDetails(address []byte) (*BalanceDetailsResponse, error) {
	addressString := crypto.Base58Encode(address)
	res := new(balanceDetailsResponse)

	path := fmt.Sprintf("/addresses/balance/details/%s", addressString)
	_, _, _, err := api.client.Get(context.Background(), path, nil, res)
	if err != nil {
		return nil, err
	}

	return &BalanceDetailsResponse{
		Address:    crypto.Base58Decode(res.Address),
		Regular:    res.Regular,
		Generating: res.Generating,
		Available:  res.Available,
		Effective:  res.Effective,
	}, nil
}
