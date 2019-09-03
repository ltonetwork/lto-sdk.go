package crypto_test

import (
	"fmt"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/stretchr/testify/require"
)

var (
	blake2bBytes = []byte{80, 159, 124, 221, 245, 42, 125, 177, 77, 112, 12, 60, 147, 143, 195, 92, 11, 59, 146, 209, 80, 78, 215, 8, 114, 147, 138, 9, 168, 93, 88, 196}
	blake2bStr   = []byte(`112233445566778899 Saturn V rocket’s first stage carries 203,400 gallons (770,000 liters) of kerosene fuel and 318,000 gallons (1.2 million liters) of liquid oxygen needed for combustion. At liftoff, the stage’s five F-1 rocket engines ignite and produce 7.5 million pounds of thrust.`)
)

func TestBlake2b(t *testing.T) {
	cases := map[string]struct {
		Input          []byte
		ExpectedOutput []byte
	}{
		"blake2b string": {
			blake2bStr,
			blake2bBytes,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result := crypto.Blake2b(tc.Input)

			require.Equal(t, tc.ExpectedOutput, result)

			fmt.Printf("output: %v\n", result)
		})
	}
}
