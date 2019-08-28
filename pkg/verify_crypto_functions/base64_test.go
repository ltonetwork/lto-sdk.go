package verify_crypto_functions_test

import (
	"fmt"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/verify_crypto_functions"
	"github.com/stretchr/testify/require"
)

var (
	b64Bytes = []byte{84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 85, 105, 110, 116,
		56, 65, 114, 114, 97, 121, 32, 99, 111, 110, 118, 101, 114, 116,
		101, 100, 32, 116, 111, 32, 97, 32, 115, 116, 114, 105, 110, 103}
	b64Str = "VGhpcyBpcyBhIFVpbnQ4QXJyYXkgY29udmVydGVkIHRvIGEgc3RyaW5n"
)

func TestBase64Encode(t *testing.T) {
	cases := map[string]struct {
		Input          []byte
		ExpectedOutput string
	}{
		"base64 encode bytes": {
			b64Bytes,
			b64Str,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := verify_crypto_functions.Base64Encode(tc.Input)

			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}

func TestBase64Decode(t *testing.T) {
	cases := map[string]struct {
		Input          string
		ExpectedOutput []byte
	}{
		"base64 decode string": {
			b64Str,
			b64Bytes,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := verify_crypto_functions.Base64Decode(tc.Input)

			require.NoError(t, err)
			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}
