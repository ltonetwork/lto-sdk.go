package lto

import (
	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/pkg/errors"
)

func init() {
	DefaultMainNetClient, _ = Client().WithNetwork(NetworkMain).New()
	DefaultTestNetClient, _ = Client().WithNetwork(NetworkTest).New()
}

var DefaultMainNetClient *client
var DefaultTestNetClient *client

type Network byte

const NetworkMain Network = 'L'
const NetworkTest Network = 'T'

var DefaultBasicConfig = &BasicConfig{
	RequestOffset:     0,
	RequestLimit:      100,
	MinimumSeedLength: 15,
	TimeDiff:          0,
}

func DefaultMainNetConfig() *Config {
	return &Config{
		BasicConfig: DefaultBasicConfig,
		Network:     NetworkMain,
		NodeAddress: "https://nodes.legalthings.one",
	}
}

func DefaultTestNetConfig() *Config {
	return &Config{
		BasicConfig: DefaultBasicConfig,
		Network:     NetworkTest,
		NodeAddress: "https://testnet.legalthings.one",
	}
}

type clientParams struct {
	config      *Config
	network     Network
	nodeAddress string
}

func Client() *clientParams {
	return &clientParams{
		network: NetworkMain,
	}
}

func (p *clientParams) New() (*client, error) {
	if p.config == nil {
		switch p.network {
		case NetworkMain:
			p.config = DefaultMainNetConfig()
		case NetworkTest:
			p.config = DefaultTestNetConfig()
		default:
			return nil, errors.New("invalid network")
		}
	}

	if p.nodeAddress != "" {
		p.config.NodeAddress = p.nodeAddress
	}

	return &client{
		Config: p.config,
	}, nil
}

func (p *clientParams) WithNodeAddress(nodeAddress string) *clientParams {
	p.nodeAddress = nodeAddress

	return p
}

func (p *clientParams) WithNetwork(network Network) *clientParams {
	p.network = network

	return p
}

func (p *clientParams) WithNetworkConfig(config *Config) *clientParams {
	p.config = config

	return p
}

type client struct {
	Config *Config
}

type Config struct {
	*BasicConfig
	Network     Network
	NodeAddress string
}

type BasicConfig struct {
	RequestOffset     int
	RequestLimit      int
	MinimumSeedLength int
	TimeDiff          int64
}

func (c *client) Account() *accountParams {
	return Account().WithNetworkConfig(c.Config)
}

func (c *client) IsValidAddress(address []byte) bool {
	return crypto.IsValidAddress(address, byte(c.Config.Network))
}

/**
 * Create an event chain id based on a public sign key
 *
 * @param publicSignKey {string} - Public sign on which the event chain will be based
 * @param nonce {string} - (optional) A random nonce will generate by default
 */
func (c *client) CreateEventChainID(publicSignKey []byte, nonce []byte) ([]byte, error) {
	chain := NewEventChain(nil)

	err := chain.Init(publicSignKey, nonce)
	if err != nil {
		return nil, err
	}

	return chain.ID, nil
}
