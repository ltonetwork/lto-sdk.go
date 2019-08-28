package verify_crypto_functions

import "crypto/sha256"

func Sha256(s string) []byte {
	hash := sha256.Sum256([]byte(s))

	return hash[:]
}
