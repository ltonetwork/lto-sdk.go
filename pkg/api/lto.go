package api

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/pkg/errors"
)

const MainNetByte byte = 'L'
const TestNetByte byte = 'T'

var DefaultBasicConfig = &LTOBasicConfig{
	RequestOffset:     0,
	RequestLimit:      100,
	LogLevel:          "warning",
	MinimumSeedLength: 15,
	TimeDiff:          0,
}

var DefaultMainNetConfig = &LTOConfig{
	LTOBasicConfig: DefaultBasicConfig,
	NetworkByte:    MainNetByte,
	NodeAddress:    "https://nodes.legalthings.one",
}
var DefaultTestNetConfig = &LTOConfig{
	LTOBasicConfig: DefaultBasicConfig,
	NetworkByte:    TestNetByte,
	NodeAddress:    "https://testnet.legalthings.one",
}

func NewLTO(networkByte byte, nodeAddress string) (*LTO, error) {
	var config *LTOConfig

	switch networkByte {
	case MainNetByte:
		config = DefaultMainNetConfig
	case TestNetByte:
		config = DefaultTestNetConfig
	default:
		return nil, errors.New("invalid network byte")
	}

	if nodeAddress != "" {
		config.NodeAddress = nodeAddress
	}

	return &LTO{
		NetworkByte: networkByte,
		Config:      config,
	}, nil
}

type LTO struct {
	NetworkByte byte
	Config      *LTOConfig
}

type LTOConfig struct {
	*LTOBasicConfig
	NetworkByte byte
	NodeAddress string
}

type LTOBasicConfig struct {
	RequestOffset     int
	RequestLimit      int
	LogLevel          string
	MinimumSeedLength int
	TimeDiff          int64
}

/**
 * Creates an account based on a random seed
 */
func (lto *LTO) CreateAccount(words int) (*Account, error) {
	if words < lto.Config.MinimumSeedLength {
		return nil, errors.Errorf("seed must have a length of at least %d", lto.Config.MinimumSeedLength)
	}

	phrase, err := generateNewSeed(words)
	if err != nil {
		return nil, err
	}

	return lto.CreateAccountFromExistingPhrase(phrase)

}

func generateNewSeed(length int) ([]byte, error) {
	dictionarySize := len(seedDictionary)
	words := make([]string, length)

	for i := 0; i < length; i++ {
		randomBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(dictionarySize)))
		if err != nil {
			return nil, err
		}

		idx := int(randomBigInt.Int64()) % dictionarySize

		words[i] = seedDictionary[idx]
	}

	return []byte(strings.Join(words, " ")), nil
}

/**
 * Creates an account based on an existing seed
 */
func (lto *LTO) CreateAccountFromExistingPhrase(phrase []byte) (*Account, error) {
	if len(phrase) < lto.Config.MinimumSeedLength {
		return nil, errors.Errorf("seed must have a length of at least %d", lto.Config.MinimumSeedLength)
	}
	//fmt.Printf("phrase = ('%v')\n", phrase)

	account, err := NewAccount(phrase, lto.NetworkByte)
	if err != nil {
		return nil, err
	}

	return account, nil
}

/**
 * Creates an account based on a private key
 */
func (lto *LTO) CreateAccountFromPrivateKey(privateKey []byte) (*Account, error) {
	account, err := NewAccount(nil, lto.NetworkByte)
	if err != nil {
		return nil, err
	}

	account.Sign = crypto.BuildNACLSignKeyPairFromSecret(privateKey)
	account.Address = crypto.BuildRawAddress(account.Sign.PublicKey, lto.Config.NetworkByte)

	return account, nil
}

func (lto *LTO) IsValidAddress(address []byte) bool {
	return crypto.IsValidAddress(address, lto.NetworkByte)
}

/**
 * Create an event chain id based on a public sign key
 *
 * @param publicSignKey {string} - Public sign on which the event chain will be based
 * @param nonce {string} - (optional) A random nonce will generate by default
 */
func (lto *LTO) CreateEventChainID(publicSignKey []byte, nonce []byte) ([]byte, error) {
	account, err := NewAccount(nil, 0)
	if err != nil {
		return nil, err
	}

	account.Sign = &crypto.KeyPair{
		PublicKey: publicSignKey,
	}

	chain, err := account.CreateEventChain(nonce)
	if err != nil {
		return nil, err
	}

	return chain.ID, nil
}
