package verify_crypto_functions

import "github.com/btcsuite/btcutil/base58"

func Base58Encode(b []byte) string {
	return base58.Encode(b)
}

func Base58Decode(s string) []byte {
	return base58.Decode(s)
}
