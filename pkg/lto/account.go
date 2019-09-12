package lto

import (
	cryptorand "crypto/rand"
	"math/big"
	"strings"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/pkg/errors"
)

type accountParams struct {
	network       Network
	networkConfig *Config
	seed          []byte
	privateKey    []byte
	randomWordN   int
}

func NewAccount() *accountParams {
	return &accountParams{
		network:       NetworkMain,
		seed:          nil,
		privateKey:    nil,
		networkConfig: nil,
		randomWordN:   15,
	}
}

func (p *accountParams) Create() (*Account, error) {
	if p.networkConfig == nil {
		switch p.network {
		case NetworkMain:
			p.networkConfig = DefaultMainNetConfig()
		case NetworkTest:
			p.networkConfig = DefaultTestNetConfig()
		default:
			return nil, errors.New("invalid network")
		}
	}

	if len(p.privateKey) != 0 {
		sign := crypto.BuildNACLSignKeyPairFromSecret(p.privateKey)
		address := crypto.BuildRawAddress(sign.PublicKey, byte(p.networkConfig.Network))

		return &Account{
			Address: address,
			Sign:    sign,
		}, nil
	}

	if len(p.seed) != 0 {
		return newAccountFromSeed(p.seed, p.networkConfig)
	}

	if p.randomWordN != 0 {
		seed, err := generateNewSeed(p.randomWordN)
		if err != nil {
			return nil, err
		}

		return newAccountFromSeed(seed, p.networkConfig)
	}

	return nil, errors.New("no method specified for generating the private key")
}

func newAccountFromSeed(seed []byte, networkConfig *Config) (*Account, error) {
	if len(seed) < networkConfig.MinimumSeedLength {
		return nil, errors.Errorf("seed must have a length of at least %d", networkConfig.MinimumSeedLength)
	}

	keys, err := crypto.BuildNACLSignKeyPair(seed)
	if err != nil {
		return nil, err
	}

	return &Account{
		Address: crypto.BuildRawAddress(keys.PublicKey, byte(networkConfig.Network)),
		Seed:    seed,
		Sign:    keys,
	}, nil
}

func (p *accountParams) FromPrivateKey(privateKey []byte) *accountParams {
	p.privateKey = privateKey

	return p
}

func (p *accountParams) FromSeed(seed []byte) *accountParams {
	p.seed = seed

	return p
}

func (p *accountParams) FromRandom() *accountParams {
	p.randomWordN = 15

	return p
}

func (p *accountParams) FromRandomN(n int) *accountParams {
	p.randomWordN = n

	return p
}

func (p *accountParams) WithNetwork(network Network) *accountParams {
	p.network = network

	return p
}

func (p *accountParams) WithNetworkConfig(config *Config) *accountParams {
	p.networkConfig = config

	return p
}

type Account struct {
	/**
	 * Client Wallet Address
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
 * Create an Event chain
 */
func (a *Account) CreateEventChain(nonce []byte) (*EventChain, error) {
	eventChain, err := NewEventChain().WithPublicKey(a.Sign.PublicKey).WithNonce(nonce).Create()
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
 * Add a signature to the Event
 */
func (a *Account) SignEvent(event *Event) (*Event, error) {
	var err error

	event.SignKey = a.Sign.PublicKey
	message, err := event.GetMessage()
	if err != nil {
		return nil, err
	}
	event.Signature, err = a.SignMessage(message)
	if err != nil {
		return nil, err
	}

	event.Hash, err = event.GetHash()
	if err != nil {
		return nil, err
	}

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

func (a *Account) GetRandomNonce() ([]byte, error) {
	bytes := make([]byte, 24)
	_, err := Rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func generateNewSeed(length int) ([]byte, error) {
	dictionarySize := len(seedDictionary)
	words := make([]string, length)

	for i := 0; i < length; i++ {
		randomBigInt, err := cryptorand.Int(Rand, big.NewInt(int64(dictionarySize)))
		if err != nil {
			return nil, err
		}

		idx := int(randomBigInt.Int64()) % dictionarySize

		words[i] = seedDictionary[idx]
	}

	return []byte(strings.Join(words, " ")), nil
}
