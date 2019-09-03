package crypto

import "golang.org/x/crypto/blake2b"

func Blake2b(s []byte) []byte {
	hash := blake2b.Sum256(s)

	return hash[:]
}
