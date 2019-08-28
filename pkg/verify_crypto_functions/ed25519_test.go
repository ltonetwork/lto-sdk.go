package verify_crypto_functions_test

import (
	"fmt"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/verify_crypto_functions"
	"github.com/stretchr/testify/require"
)

var (
	signature  = []byte{52, 196, 32, 121, 138, 28, 241, 111, 11, 191, 80, 17, 255, 225, 161, 117, 55, 217, 56, 207, 131, 177, 177, 109, 219, 203, 101, 12, 199, 151, 0, 104, 121, 66, 87, 206, 113, 47, 84, 210, 140, 44, 41, 24, 91, 148, 39, 149, 46, 219, 76, 255, 224, 28, 104, 233, 25, 80, 138, 27, 133, 123, 89, 0}
	publicKey  = []byte{89, 250, 0, 83, 37, 125, 48, 176, 137, 212, 169, 165, 253, 188, 41, 231, 10, 50, 123, 71, 15, 116, 225, 233, 199, 150, 201, 84, 168, 114, 132, 127}
	privateKey = []byte{95, 5, 202, 39, 42, 215, 250, 254, 100, 4, 234, 129, 190, 28, 91, 250, 197, 112, 169, 190, 47, 67, 133, 92, 172, 193, 201, 200, 164, 166, 234, 194, 89, 250, 0, 83, 37, 125, 48, 176, 137, 212, 169, 165, 253, 188, 41, 231, 10, 50, 123, 71, 15, 116, 225, 233, 199, 150, 201, 84, 168, 114, 132, 127}
	message    = []byte{84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 85, 105, 110, 116, 56, 65, 114, 114, 97, 121, 32, 99, 111, 110, 118, 101, 114, 116, 101, 100, 32, 116, 111, 32, 97, 32, 115, 116, 114, 105, 110, 103}
)

func TestED25519Sign(t *testing.T) {
	cases := map[string]struct {
		PrivateKey     []byte
		Message        []byte
		ExpectedOutput []byte
	}{
		"ed25519 sign message": {
			privateKey,
			message,
			signature,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := verify_crypto_functions.ED25519Sign(tc.PrivateKey, tc.Message)

			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}

func TestED25519Verify(t *testing.T) {
	cases := map[string]struct {
		PublicKey      []byte
		Message        []byte
		Signature      []byte
		ExpectedOutput bool
	}{
		"ed25519 verify message": {
			publicKey,
			message,
			signature,
			true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := verify_crypto_functions.ED25519Verify(tc.PublicKey, tc.Message, tc.Signature)

			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}
