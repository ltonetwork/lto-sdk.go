package xapi

import (
	"bytes"
	"encoding/json"
	"io"
)

type JSONEncoder struct {
}

func NewJSONEncoder() *JSONEncoder {
	return &JSONEncoder{}
}

func (*JSONEncoder) Encode(body interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(body)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (*JSONEncoder) Decode(body io.Reader, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}

func (*JSONEncoder) Type() string {
	return "json"
}
