package crypto

import "crypto/sha256"

func Sha256(s []byte) []byte {
	hash := sha256.Sum256(s)

	return hash[:]
}
