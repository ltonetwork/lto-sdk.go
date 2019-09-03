package crypto_test

import (
	"fmt"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/stretchr/testify/require"
)

var (
	sha256Bytes = []byte{176, 189, 180, 212, 61, 57, 135, 83, 77, 129, 113, 121, 98, 174, 108, 84, 94, 242, 230, 210, 41, 15, 167, 31, 88, 176, 183, 225, 210, 120, 246, 34}
	sha256Str   = []byte(`112233445566778899 Saturn V rocket’s first stage carries 203,400 gallons (770,000 liters) of kerosene fuel and 318,000 gallons (1.2 million liters) of liquid oxygen needed for combustion. At liftoff, the stage’s five F-1 rocket engines ignite and produce 7.5 million pounds of thrust.`)
)

func TestSha256(t *testing.T) {
	cases := map[string]struct {
		Input          []byte
		ExpectedOutput []byte
	}{
		"sha256 string": {
			sha256Str,
			sha256Bytes,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := crypto.Sha256(tc.Input)

			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}
