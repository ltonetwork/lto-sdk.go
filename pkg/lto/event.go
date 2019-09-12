package lto

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
	"github.com/pkg/errors"
)

const TimeFormat = "2006-01-02T15:04:05-07:00"

type eventParams struct {
	body         interface{}
	previousHash string
	signature    []byte
	timestamp    int64
	signKey      []byte
}

func NewEvent() *eventParams {
	return &eventParams{
		timestamp: time.Now().Unix(),
	}
}

func (p *eventParams) Create() (*Event, error) {
	var body string

	if p.body != nil {
		bodyBytes, err := json.Marshal(p.body)
		if err != nil {
			return nil, err
		}

		body = crypto.Base58Encode(bodyBytes)
	}

	return &Event{
		Timestamp: p.timestamp,
		Previous:  p.previousHash,
		Body:      body,
		Signature: p.signature,
		SignKey:   p.signKey,
	}, nil
}

func (p *eventParams) WithBody(body interface{}) *eventParams {
	p.body = body
	return p
}

func (p *eventParams) WithPrevious(previousHash string) *eventParams {
	p.previousHash = previousHash
	return p
}

func (p *eventParams) WithSignature(signature []byte) *eventParams {
	p.signature = signature
	return p
}

func (p *eventParams) WithSignKey(signKey []byte) *eventParams {
	p.signKey = signKey
	return p
}

func (p *eventParams) WithTimestamp(timestamp int64) *eventParams {
	p.timestamp = timestamp
	return p
}

type Event struct {
	/**
	 * Base58 encoded JSON string with the body of the Event.
	 *
	 */
	Body string

	/**
	 * Time when the Event was signed.
	 *
	 */
	Timestamp int64

	/**
	 * Hash to the previous Event
	 *
	 */
	Previous string

	/**
	 * URI of the public key used to sign the Event
	 *
	 */
	SignKey []byte

	/**
	 * Base58 encoded signature of the Event
	 *
	 */
	Signature []byte

	/**
	 * Base58 encoded SHA256 hash of the Event
	 *
	 */
	Hash string
}

func (e *Event) GetHash() (string, error) {
	message, err := e.GetMessage()
	if err != nil {
		return "", err
	}
	return crypto.Base58Encode(crypto.Sha256(message)), nil
}

func (e *Event) GetMessage() ([]byte, error) {
	if len(e.Body) == 0 {
		return nil, errors.New("body unknown")
	}
	if len(e.SignKey) == 0 {
		return nil, errors.New("first set signkey before creating message")
	}

	return []byte(fmt.Sprintf("%s\n%d\n%s\n%s", e.Body, e.Timestamp, e.Previous, crypto.Base58Encode(e.SignKey))), nil
}

func (e *Event) GetResourceVersion() string {
	return crypto.Base58Encode(crypto.Sha256([]byte(e.Body)))[0:8]
}

func (e *Event) VerifySignature() (bool, error) {
	message, err := e.GetMessage()
	if err != nil {
		return false, err
	}

	return crypto.VerifySignature(message, e.Signature, e.SignKey)
}

func (e *Event) GetBody(obj interface{}) error {
	err := json.Unmarshal(crypto.Base58Decode(e.Body), &obj)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) SignWith(account *Account) (*Event, error) {
	return account.SignEvent(e)
}

func (e *Event) AddTo(chain *EventChain) (*Event, error) {
	return chain.AddEvent(e)
}
