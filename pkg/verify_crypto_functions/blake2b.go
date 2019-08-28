package verify_crypto_functions

import "golang.org/x/crypto/blake2b"

func Blake2b(s string) []byte {
	hash := blake2b.Sum256([]byte(s))

	return hash[:]
}
