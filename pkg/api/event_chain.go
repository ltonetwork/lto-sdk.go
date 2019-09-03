package api

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
)
import "crypto/rand"

func NewEventChain(id string) *EventChain {
	return &EventChain{
		ID: id,
	}
}

type EventChain struct {
	ID     string
	Events []*Event
}

const EventChainVersion byte = 0x40
const ProjectionAddressVersion byte = 0x50

func (e *EventChain) Init(account *Account, nonce []byte) error {
	nonceBytes, err := e.getNonceBytes(nonce)
	if err != nil {
		return err
	}

	e.ID = crypto.BuildEventChainID(EventChainVersion, account.Sign.PublicKey, nonceBytes)
	return nil
}

func (e *EventChain) CreateProjectionID(nonce []byte) (string, error) {
	nonceBytes, err := e.getNonceBytes(nonce)
	if err != nil {
		return "", err
	}

	return crypto.BuildEventChainID(ProjectionAddressVersion, []byte(e.ID), nonceBytes), nil
}

func (e *EventChain) AddEvent(event *Event) *Event {
	event.Previous = e.GetLatestHash()
	e.Events = append(e.Events, event)

	return event
}

func (e *EventChain) GetLatestHash() string {
	if len(e.Events) == 0 {
		return crypto.BuildHash(base58.Decode(e.ID))
	}

	event := e.Events[len(e.Events)-1]
	return event.GetHash()
}

func (e *EventChain) getNonceBytes(nonce []byte) ([]byte, error) {
	var err error
	var nonceBytes []byte

	if len(nonce) == 0 {
		nonceBytes, err = e.GetRandomNonce()
	} else {
		nonceBytes = e.createNonce(nonce)
	}
	if err != nil {
		return nil, err
	}

	return nonceBytes, nil
}

func (e *EventChain) createNonce(input []byte) []byte {
	return crypto.Sha256(input)[0:20]
}

func (e *EventChain) GetRandomNonce() ([]byte, error) {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
