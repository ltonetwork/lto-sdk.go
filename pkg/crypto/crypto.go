package crypto

import (
	"github.com/pkg/errors"

	"golang.org/x/crypto/ed25519"
)

func BuildNACLSignKeyPair(seed []byte) (*KeyPair, error) {
	seedHash, err := buildSeedHash(seed)
	if err != nil {
		return nil, err
	}

	privateKey := ed25519.NewKeyFromSeed(seedHash)

	return &KeyPair{
		PublicKey:  privateKey,
		PrivateKey: privateKey.Public().(ed25519.PublicKey),
	}, nil
}

const InitialNonce = 0
const AddressVersion byte = 0x1
const PrivateKeyLength = 64
const PublicKeyLength = 32
const SignatureLength = 64

func buildSeedHash(seed []byte) ([]byte, error) {
	seedBytesWithNonce := append(seed, InitialNonce)
	seedHash := hashChain(seedBytesWithNonce)
	return Sha256(seedHash), nil
}

func hashChain(input []byte) []byte {
	return Sha256(Blake2b(input))
}

func BuildRawAddress(publicKeyBytes []byte, networkByte byte) string {
	prefix := []byte{AddressVersion, networkByte}
	publicKeyHashPart := hashChain(publicKeyBytes)[0:20]

	rawAddress := append(prefix, publicKeyHashPart...)
	addressHash := hashChain(rawAddress)[0:4]

	return Base58Encode(append(rawAddress, addressHash...))
}

func CreateSignature(input []byte, privateKey []byte) ([]byte, error) {
	if len(privateKey) != PrivateKeyLength {
		return nil, errors.New("invalid private key")
	}

	return ED25519Sign(privateKey, input), nil
}

func VerifySignature(input []byte, signature []byte, publicKey []byte) (bool, error) {
	if len(publicKey) != PublicKeyLength {
		return false, errors.New("invalid public key")
	}

	if len(signature) != SignatureLength {
		return false, errors.New("invalid signature size")
	}

	return ED25519Verify(publicKey, input, signature), nil
}

func BuildEventChainID(prefix byte, publicKey []byte, randomBytes []byte) string {
	publicKeyHashPart := hashChain(publicKey)
	rawID := append([]byte{prefix}, randomBytes...)
	rawID = append(rawID, publicKeyHashPart...)
	addressHash := hashChain(rawID)[0:4]

	return Base58Encode(append(rawID, addressHash...))
}

func BuildHash(eventBytes []byte) string {
	return Base58Encode(Sha256(eventBytes))
}
