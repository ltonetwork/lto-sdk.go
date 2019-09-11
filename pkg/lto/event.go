package lto

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"
)

const TimeFormat = "2006-01-02T15:04:05-07:00"

//const TimeFormat = time.RFC3339

func NewEvent(body interface{}, previous string) (*Event, error) {
	event := &Event{
		Timestamp: time.Now().Unix(),
		Previous:  previous,
	}

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		event.Body = crypto.Base58Encode(bodyBytes)
	}

	return event, nil
}

type Event struct {
	/**
	 * Base58 encoded JSON string with the body of the event.
	 *
	 */
	Body string

	/**
	 * Time when the event was signed.
	 *
	 */
	Timestamp int64

	/**
	 * Hash to the previous event
	 *
	 */
	Previous string

	/**
	 * URI of the public key used to sign the event
	 *
	 */
	SignKey []byte

	/**
	 * Base58 encoded signature of the event
	 *
	 */
	Signature []byte

	/**
	 * Base58 encoded SHA256 hash of the event
	 *
	 */
	Hash string
}

func (e *Event) GetHash() string {
	return crypto.Base58Encode(crypto.Sha256(e.GetMessage()))
}

func (e *Event) GetMessage() []byte {
	return []byte(fmt.Sprintf("%s\n%d\n%s\n%s", e.Body, e.Timestamp, e.Previous, crypto.Base58Encode(e.SignKey)))
}

func (e *Event) GetResourceVersion() string {
	return crypto.Base58Encode(crypto.Sha256([]byte(e.Body)))[0:8]
}

func (e *Event) VerifySignature() (bool, error) {
	return crypto.VerifySignature(e.GetMessage(), e.Signature, e.SignKey)
}

func (e *Event) GetBody(obj interface{}) error {
	err := json.Unmarshal(crypto.Base58Decode(e.Body), &obj)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) SignWith(account *account) (*Event, error) {
	return account.SignEvent(e)
}

func (e *Event) AddTo(chain *EventChain) *Event {
	return chain.AddEvent(e)
}
