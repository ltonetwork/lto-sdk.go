package verify_crypto_functions_test

import (
	"fmt"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/verify_crypto_functions"
	"github.com/stretchr/testify/require"
)

var (
	b58Bytes = []byte{84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 85, 105, 110, 116,
		56, 65, 114, 114, 97, 121, 32, 99, 111, 110, 118, 101, 114, 116,
		101, 100, 32, 116, 111, 32, 97, 32, 115, 116, 114, 105, 110, 103}
	b58Str = "2QhwuRSWNgihbhyZmj9aMe1qpoVhHs6KyXqu8XAA2dthicp2G9uY7DnTGJ"
)

func TestBase58Encode(t *testing.T) {
	cases := map[string]struct {
		Input          []byte
		ExpectedOutput string
	}{
		"base58 encode bytes": {
			b58Bytes,
			b58Str,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := verify_crypto_functions.Base58Encode(tc.Input)

			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}

func TestBase58Decode(t *testing.T) {
	cases := map[string]struct {
		Input          string
		ExpectedOutput []byte
	}{
		"base58 decode string": {
			b58Str,
			b58Bytes,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := verify_crypto_functions.Base58Decode(tc.Input)

			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}
