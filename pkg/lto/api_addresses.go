package lto

import (
	"fmt"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/pkg/errors"
)

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

func (api *API) AddressBalance(address []byte) (*BalanceResponse, error) {
	addressString := crypto.Base58Encode(address)
	res := new(balanceResponse)

	path := fmt.Sprintf("/addresses/balance/%s", addressString)
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get balance")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return &BalanceResponse{
		Address:       crypto.Base58Decode(res.Address),
		Confirmations: res.Confirmations,
		Balance:       res.Balance,
	}, nil
}

func (api *API) AddressBalanceWithConfirmations(address []byte, confirmations int) (*BalanceResponse, error) {
	addressString := crypto.Base58Encode(address)
	res := new(balanceResponse)

	path := fmt.Sprintf("/addresses/balance/%s/%d", addressString, confirmations)
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get balance")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
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

func (api *API) AddressBalanceDetails(address []byte) (*BalanceDetailsResponse, error) {
	addressString := crypto.Base58Encode(address)
	res := new(balanceDetailsResponse)

	path := fmt.Sprintf("/addresses/balance/details/%s", addressString)
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get balance details")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return &BalanceDetailsResponse{
		Address:    crypto.Base58Decode(res.Address),
		Regular:    res.Regular,
		Generating: res.Generating,
		Available:  res.Available,
		Effective:  res.Effective,
	}, nil
}
