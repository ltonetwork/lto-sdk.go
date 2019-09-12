package lto

import (
	"errors"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
)

type eventChainParams struct {
	id        []byte
	publicKey []byte
	nonce     []byte
}

func NewEventChain() *eventChainParams {
	return &eventChainParams{}
}

func (p *eventChainParams) Create() (*EventChain, error) {
	var id []byte

	if p.id != nil {
		id = p.id
	} else if p.publicKey != nil {
		nonceBytes, err := getNonceBytes(p.nonce)
		if err != nil {
			return nil, err
		}

		id = crypto.BuildEventChainID(EventChainVersion, p.publicKey, nonceBytes)
	}

	return &EventChain{
		ID: id,
	}, nil
}

func (p *eventChainParams) WithID(id []byte) *eventChainParams {
	p.id = id
	return p
}

func (p *eventChainParams) WithPublicKey(publicKey []byte) *eventChainParams {
	p.publicKey = publicKey
	return p
}

func (p *eventChainParams) WithNonce(nonce []byte) *eventChainParams {
	p.nonce = nonce
	return p
}

type EventChain struct {
	ID     []byte
	Events []*Event
}

const EventChainVersion byte = 0x40
const ProjectionAddressVersion byte = 0x50

func (e *EventChain) CreateProjectionID(nonce []byte) ([]byte, error) {
	if len(e.ID) == 0 {
		return nil, errors.New("no id set for projection id")
	}

	nonceBytes, err := getNonceBytes(nonce)
	if err != nil {
		return nil, err
	}

	return crypto.BuildEventChainID(ProjectionAddressVersion, []byte(crypto.Base58Encode(e.ID)), nonceBytes), nil
}

func (e *EventChain) AddEvent(event *Event) (*Event, error) {
	var err error

	event.Previous, err = e.GetLatestHash()
	if err != nil {
		return nil, err
	}

	e.Events = append(e.Events, event)

	return event, nil
}

func (e *EventChain) GetLatestHash() (string, error) {
	if len(e.Events) == 0 {
		return crypto.BuildHash(e.ID), nil
	}

	event := e.Events[len(e.Events)-1]

	return event.GetHash()
}

func getNonceBytes(nonce []byte) ([]byte, error) {
	var err error
	var nonceBytes []byte

	if len(nonce) == 0 {
		nonceBytes, err = GetRandomNonce()
	} else {
		nonceBytes = createNonce(nonce)
	}
	if err != nil {
		return nil, err
	}

	return nonceBytes, nil
}

func createNonce(input []byte) []byte {
	return crypto.Sha256(input)[0:20]
}

func GetRandomNonce() ([]byte, error) {
	bytes := make([]byte, 20)
	_, err := Rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
