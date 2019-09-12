package lto

import (
	cryptorand "crypto/rand"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/pkg/errors"
)

func init() {
	DefaultMainNetClient, _ = NewClient().WithNetwork(NetworkMain).Create()
	DefaultTestNetClient, _ = NewClient().WithNetwork(NetworkTest).Create()
}

var Rand = cryptorand.Reader

var DefaultMainNetClient *Client
var DefaultTestNetClient *Client

type Network byte

const NetworkMain Network = 'L'
const NetworkTest Network = 'T'

func DefaultBasicConfig() *BasicConfig {
	return &BasicConfig{
		RequestOffset:     0,
		RequestLimit:      100,
		MinimumSeedLength: 15,
		TimeDiff:          0,
	}
}

func DefaultMainNetConfig() *Config {
	return &Config{
		BasicConfig: DefaultBasicConfig(),
		Network:     NetworkMain,
		NodeAddress: "https://nodes.legalthings.one",
	}
}

func DefaultTestNetConfig() *Config {
	return &Config{
		BasicConfig: DefaultBasicConfig(),
		Network:     NetworkTest,
		NodeAddress: "https://testnet.legalthings.one",
	}
}

type clientParams struct {
	config      *Config
	network     Network
	nodeAddress string
}

func NewClient() *clientParams {
	return &clientParams{
		network: NetworkMain,
	}
}

func (p *clientParams) Create() (*Client, error) {
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

	api, err := NewAPI(p.config)
	if err != nil {
		return nil, err
	}

	return &Client{
		Config: p.config,
		API:    api,
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

type Client struct {
	Config *Config
	*API
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

func (c *Client) NewAccount() *accountParams {
	return NewAccount().WithNetworkConfig(c.Config)
}

func (c *Client) IsValidAddress(address []byte) bool {
	return crypto.IsValidAddress(address, byte(c.Config.Network))
}

/**
 * Create an Event chain id based on a public sign key
 *
 * @param publicSignKey {string} - Public sign on which the Event chain will be based
 * @param nonce {string} - (optional) A random nonce will generate by default
 */
func (c *Client) CreateEventChainID(publicSignKey []byte, nonce []byte) ([]byte, error) {
	chain, err := NewEventChain().WithPublicKey(publicSignKey).WithNonce(nonce).Create()
	if err != nil {
		return nil, err
	}

	return chain.ID, nil
}
