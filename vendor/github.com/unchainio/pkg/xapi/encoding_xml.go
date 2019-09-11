package xapi

import (
	"bytes"
	"encoding/xml"
	"io"
)

type XMLEncoder struct {
}

func NewXMLEncoder() *XMLEncoder {
	return &XMLEncoder{}
}

func (*XMLEncoder) Encode(body interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	enc := xml.NewEncoder(buf)
	err := enc.Encode(body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (*XMLEncoder) Decode(body io.Reader, v interface{}) error {
	return xml.NewDecoder(body).Decode(v)
}

func (*XMLEncoder) Type() string {
	return "xml"
}
