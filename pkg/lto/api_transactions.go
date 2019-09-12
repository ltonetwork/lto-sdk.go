package lto

import (
	"fmt"

	"github.com/pkg/errors"
)

type TransactionsGetResponse struct {
	Type      int64  `json:"type"`
	ID        string `json:"id"`
	Fee       int64  `json:"fee"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
	Recipient string `json:"recipient"`
	Amount    int64  `json:"amount"`
	Height    int64  `json:"height"`
}

func (api *API) TransactionsGet(id string) (*TransactionsGetResponse, error) {
	res := new(TransactionsGetResponse)

	path := fmt.Sprintf("/transactions/info/%s", id)
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", string(r.Body()))

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

type TransactionsGetListResponseItem struct {
	Type            int64  `json:"type"`
	ID              string `json:"id"`
	Sender          string `json:"sender"`
	SenderPublicKey string `json:"senderPublicKey"`
	Fee             int64  `json:"fee"`
	Timestamp       int64  `json:"timestamp"`
	Signature       string `json:"signature"`
	Version         int64  `json:"version"`
	Recipient       string `json:"recipient"`
	Amount          int64  `json:"amount"`
	Status          string `json:"status"`
	Attachment      string `json:"attachment"`
	Height          int64  `json:"height"`
}

func (api *API) TransactionsGetList(address string, limit int) ([][]*TransactionsGetListResponseItem, error) {
	if limit == 0 {
		limit = api.config.RequestLimit
	}

	var res [][]*TransactionsGetListResponseItem

	path := fmt.Sprintf("/transactions/address/%s/limit/%d", address, limit)
	r, err := api.client.R().SetResult(&res).Get(path)
	if err != nil {
		return nil, err
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

type TransactionsUTXSizeResponse struct {
	Size int64 `json:"size"`
}

func (api *API) TransactionsUTXSize() (*TransactionsUTXSizeResponse, error) {
	var res *TransactionsUTXSizeResponse

	path := fmt.Sprintf("/transactions/unconfirmed/size")
	r, err := api.client.R().SetResult(&res).Get(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", string(r.Body()))

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

func (api *API) TransactionsUTXGet(id string) (*TransactionsGetResponse, error) {
	var res *TransactionsGetResponse

	path := fmt.Sprintf("/transactions/unconfirmed/info/%s", id)
	r, err := api.client.R().SetResult(&res).Get(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", string(r.Body()))

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

func (api *API) TransactionsUTXGetList() ([][]*TransactionsGetListResponseItem, error) {
	var res [][]*TransactionsGetListResponseItem

	path := fmt.Sprintf("/transactions/unconfirmed")
	r, err := api.client.R().SetResult(&res).Get(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", string(r.Body()))

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}
