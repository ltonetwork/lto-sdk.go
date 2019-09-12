package lto

import (
	"fmt"

	"github.com/pkg/errors"
)

type BlocksGetResponse struct {
	Version          int64           `json:"version"`
	Timestamp        int64           `json:"timestamp"`
	Reference        string          `json:"reference"`
	NXTConsensus     *NXTConsensus   `json:"nxt-consensus"`
	Generator        string          `json:"generator"`
	Signature        string          `json:"signature"`
	BlockSize        int64           `json:"blocksize"`
	TransactionCount int64           `json:"transactionCount"`
	Fee              int64           `json:"fee"`
	Transactions     []*Transactions `json:"transactions"`
	Height           int64           `json:"height"`
}

type NXTConsensus struct {
	BaseTarget          int64  `json:"base-target"`
	GenerationSignature string `json:"generation-signature"`
}

type Transactions struct {
	Type      int64  `json:"type"`
	ID        string `json:"id"`
	Fee       int64  `json:"fee"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
	Recipient string `json:"recipient"`
	Amount    int64  `json:"amount"`
}

func (api *API) BlocksGet(signature string) (*BlocksGetResponse, error) {
	res := new(BlocksGetResponse)

	path := fmt.Sprintf("/blocks/signature/%s", signature)
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

func (api *API) BlocksAt(height int64) (*BlocksGetResponse, error) {
	res := new(BlocksGetResponse)

	path := fmt.Sprintf("/blocks/at/%d", height)
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

func (api *API) BlocksFirst() (*BlocksGetResponse, error) {
	res := new(BlocksGetResponse)

	path := fmt.Sprintf("/blocks/first")
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

func (api *API) BlocksLast() (*BlocksGetResponse, error) {
	res := new(BlocksGetResponse)

	path := fmt.Sprintf("/blocks/last")
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

type BlocksHeightResponse struct {
	Height int64 `json:"height"`
}

func (api *API) BlocksHeight() (*BlocksHeightResponse, error) {
	res := new(BlocksHeightResponse)

	path := fmt.Sprintf("/blocks/height")
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}
