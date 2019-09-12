package lto

import (
	"errors"
	"fmt"
)

type UtilsTimeResponse struct {
	System int64 `json:"system"`
	NTP    int64 `json:"NTP"`
}

func (api *API) UtilsTime() (*UtilsTimeResponse, error) {
	res := new(UtilsTimeResponse)

	path := fmt.Sprintf("/utils/time")
	r, err := api.client.R().SetResult(res).Get(path)
	if err != nil {
		return nil, err
	}

	if r.IsError() {
		return nil, errors.New(string(r.Body()))
	}

	return res, nil
}

type UtilsCompileResponse struct {
	Script string `json:"script"`
}

type UtilsCompileResponseError struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func (api *API) UtilsCompile(code string) (string, error) {
	res := new(UtilsCompileResponse)

	path := fmt.Sprintf("/utils/script/compile")
	r, err := api.client.R().SetBody(code).SetResult(res).Post(path)

	if err != nil {
		return "", err
	}

	if r.IsError() {
		return "", errors.New(string(r.Body()))
	}

	return res.Script, nil
}
