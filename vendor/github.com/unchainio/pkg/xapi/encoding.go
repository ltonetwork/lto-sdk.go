package xapi

import (
	"io"
)

type Encoder interface {
	Encode(interface{}) (io.Reader, error)
	Decode(io.Reader, interface{}) error
	Type() string
}
