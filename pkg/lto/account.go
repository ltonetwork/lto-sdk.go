package lto

import (
	cryptorand "crypto/rand"
	"math/big"
	"strings"

	"github.com/davecgh/go-spew/spew"

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

func Account() *accountParams {
	return &accountParams{
		network:       NetworkMain,
		seed:          nil,
		privateKey:    nil,
		networkConfig: nil,
		randomWordN:   15,
	}
}

func (p *accountParams) New() (*account, error) {
	spew.Dump(p.randomWordN)
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

		return &account{
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

	return nil, nil
}

func newAccountFromSeed(seed []byte, networkConfig *Config) (*account, error) {
	if len(seed) < networkConfig.MinimumSeedLength {
		return nil, errors.Errorf("seed must have a length of at least %d", networkConfig.MinimumSeedLength)
	}

	keys, err := crypto.BuildNACLSignKeyPair(seed)
	if err != nil {
		return nil, err
	}

	return &account{
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

	spew.Dump(p.randomWordN)

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

type account struct {
	/**
	 * client Wallet Address
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
func (a *account) CreateEventChain(nonce []byte) (*EventChain, error) {
	eventChain := NewEventChain(nil)
	err := eventChain.Init(a.Sign.PublicKey, nonce)
	if err != nil {
		return nil, err
	}

	return eventChain, nil
}

/**
 * Get encoded seed phrase
 */
func (a *account) GetEncodedPhrase() string {
	return crypto.Base58Encode(a.Seed)
}

/**
 * Add a signature to the event
 */
func (a *account) SignEvent(event *Event) (*Event, error) {
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
func (a *account) Verify(signature []byte, message []byte) (bool, error) {
	return crypto.VerifySignature(message, signature, a.Sign.PublicKey)
}

/**
 * Create a signature from a message
 */
func (a *account) SignMessage(message []byte) ([]byte, error) {
	return crypto.CreateSignature(message, a.Sign.PrivateKey)
}

func (a *account) GetNonce() ([]byte, error) {
	bytes := make([]byte, 24)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
func generateNewSeed(length int) ([]byte, error) {
	dictionarySize := len(seedDictionary)
	words := make([]string, length)

	for i := 0; i < length; i++ {
		randomBigInt, err := cryptorand.Int(rand, big.NewInt(int64(dictionarySize)))
		if err != nil {
			return nil, err
		}

		idx := int(randomBigInt.Int64()) % dictionarySize

		words[i] = seedDictionary[idx]
	}

	return []byte(strings.Join(words, " ")), nil
}
