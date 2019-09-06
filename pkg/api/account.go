package api

import (
	"crypto/rand"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
)

func NewAccount(phrase []byte, networkByte byte) (*Account, error) {
	if len(phrase) != 0 {
		keys, err := crypto.BuildNACLSignKeyPair(phrase)
		if err != nil {
			return nil, err
		}
		//
		//fmt.Printf("privateKey = ('%v')\n", keys.PrivateKey)
		//fmt.Printf("publicKey = ('%v')\n", keys.PublicKey)

		return &Account{
			Address: crypto.BuildRawAddress(keys.PublicKey, networkByte),
			Seed:    phrase,
			Sign:    keys,
		}, nil
	}

	return &Account{}, nil
}

type Account struct {
	/**
	 * LTO Wallet Address
	 */
	Address []byte

	/**
	 * Seed phrase
	 */
	Seed []byte

	/**
	 * Signing keys
	 */
	Sign *crypto.KeyPair
}

/**
 * Create an event chain
 */
func (a *Account) CreateEventChain(nonce []byte) (*EventChain, error) {
	eventChain := NewEventChain(nil)
	err := eventChain.Init(a, nonce)
	if err != nil {
		return nil, err
	}

	return eventChain, nil
}

/**
 * Get encoded seed phrase
 */
func (a *Account) GetEncodedPhrase() string {
	return crypto.Base58Encode(a.Seed)
}

/**
 * Add a signature to the event
 */
func (a *Account) SignEvent(event *Event) (*Event, error) {
	var err error

	event.SignKey = a.Sign.PublicKey
	event.Signature, err = a.SignMessage(event.GetMessage())
	if err != nil {
		return nil, err
	}

	event.Hash = event.GetHash()
	return event, nil
}

/**
 * Verify a signature with a message
 */
func (a *Account) Verify(signature []byte, message []byte) (bool, error) {
	return crypto.VerifySignature(message, signature, a.Sign.PublicKey)
}

/**
 * Create a signature from a message
 */
func (a *Account) SignMessage(message []byte) ([]byte, error) {
	return crypto.CreateSignature(message, a.Sign.PrivateKey)
}

func (a *Account) GetNonce() ([]byte, error) {
	bytes := make([]byte, 24)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
