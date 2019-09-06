package crypto

import (
	"bytes"
	"encoding/binary"

	"github.com/pkg/errors"

	"golang.org/x/crypto/ed25519"
)

func BuildNACLSignKeyPair(seed []byte) (*KeyPair, error) {
	//fmt.Printf("seedBytes = ('%v')\n", seed)

	seedHash, err := buildSeedHash(seed)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("seedHash = ('%v')\n", seedHash)

	privateKey := ed25519.NewKeyFromSeed(seedHash)

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  privateKey.Public().(ed25519.PublicKey),
	}, nil
}

const InitialNonce int32 = 0
const AddressVersion byte = 0x1
const PrivateKeyLength = 64
const PublicKeyLength = 32
const SignatureLength = 64

func buildSeedHash(seed []byte) ([]byte, error) {
	nonce := new(bytes.Buffer)

	err := binary.Write(nonce, binary.BigEndian, InitialNonce)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("nonce = ('%v')\n", nonce.Bytes())

	seedBytesWithNonce := append(nonce.Bytes(), seed...)
	//fmt.Printf("seedBytesWithNonce = ('%v')\n", seedBytesWithNonce)

	seedHash := hashChain(seedBytesWithNonce)
	return Sha256(seedHash), nil
}

func hashChain(input []byte) []byte {
	res := Sha256(Blake2b(input))

	return res
}

func BuildRawAddress(publicKeyBytes []byte, networkByte byte) []byte {
	prefix := []byte{AddressVersion, networkByte}
	publicKeyHashPart := hashChain(publicKeyBytes)
	//fmt.Printf("hashChain('%v') = '%v'\n", Base58Encode(publicKeyBytes), Base58Encode(publicKeyHashPart))
	publicKeyHashPart = publicKeyHashPart[0:20]
	rawAddress := append(prefix, publicKeyHashPart...)
	//fmt.Printf("publicKeyBytes = ('%v')\n", publicKeyBytes)
	//fmt.Printf("hashChain(publicKeyBytes) = ('%v')\n", hashChain(publicKeyBytes))
	//fmt.Printf("prefix = ('%v')\n", prefix)
	//fmt.Printf("publicKeyHashPart = ('%v')\n", publicKeyHashPart)
	//fmt.Printf("rawAddress = ('%v')\n", rawAddress)
	addressHash := hashChain(rawAddress)[0:4]

	return append(rawAddress, addressHash...)
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

func BuildEventChainID(prefix byte, publicKey []byte, randomBytes []byte) []byte {
	publicKeyHashPart := hashChain(publicKey)[0:20]
	rawID := append([]byte{prefix}, randomBytes...)
	rawID = append(rawID, publicKeyHashPart...)
	addressHash := hashChain(rawID)[0:4]

	return append(rawID, addressHash...)
}

func BuildHash(eventBytes []byte) string {
	return Base58Encode(Sha256(eventBytes))
}

func BuildNACLSignKeyPairFromSecret(privateKey []byte) *KeyPair {
	return &KeyPair{
		PublicKey:  ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey),
		PrivateKey: privateKey,
	}
}

func IsValidAddress(address []byte, networkByte byte) bool {
	if len(address) < 2 || address[0] != AddressVersion || address[1] != networkByte {
		return false
	}

	key := address[0:22]
	check := address[22:26]
	keyHash := hashChain(key)[0:4]

	for i := 0; i < 4; i++ {
		if check[i] != keyHash[i] {
			return false
		}
	}

	return true
}
