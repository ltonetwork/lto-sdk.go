# LTO Network client for Golang
Signing and addresses work both for the (private) event chain as for the public chain.

## Installation
```sh
go get github.com/ltonetwork/lto-sdk.go
```
## Usage
```go
package main
import "github.com/ltonetwork/lto-sdk.go/pkg/lto"
func main() {
account := lto.NewAccount()
}
```

## Accounts

### Creation

#### Create an account from seed
```go
seed := []byte("manage manual recall harvest series desert melt police rose hollow moral pledge kitten position add")
	account, err := lto.NewAccount().FromSeed(seed).Create()
	if err != nil {
		log.Fatal(err)
	}
	accountInfo := account.GetEncodedPhrase()
```

#### Create an account from sign key

```go
Sign := &crypto.KeyPair{
		PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
		PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
	}
	account, err := lto.NewAccount().FromPrivateKey(Sign.PrivateKey).Create()
```

## Signing
### Sign a message
```go
Sign := &crypto.KeyPair{
		PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
		PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
	}
	account, err := lto.NewAccount().FromPrivateKey(Sign.PrivateKey).Create()
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("hello")
	signedMessage, err := account.SignMessage(message)
	if err != nil {
		log.Error("Error Occurred While Signing Message: ", err)
	}
    fmt.Println(string(signedMessage))
```
### Create signature for an Event

```go
time1, err := time.Parse(lto.TimeFormat, "2018-03-01T00:00:00+00:00")

	Sign := &crypto.KeyPair{
		PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
		PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
	}
	account, err := lto.NewAccount().FromPrivateKey(Sign.PrivateKey).Create()
	if err != nil {
		log.Fatal("Not able to create account", err)
	}

	event := &lto.Event{
		Body:      "HeFMDcuveZQYtBePVUugLyWtsiwsW4xp7xKdv",
		Timestamp: time1.Unix(),
		Previous:  "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
		SignKey:   crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
	}
	signedEvent, err := account.SignEvent(event)
	if err != nil {
		log.Error("SignEvent() error = %v", err)
	}
	fmt.Println(string(signedEvent.Hash))
	fmt.Println(string(signedEvent.Signature))
```
### Verify a signature

```go
isValid, err := account.Verify(signedMessage, "hello")
```

## Event Chain
### Create a new event chain
```go
	nonce := []byte("10")
	chain, err := account.CreateEventChain(nonce)
	if err != nil {
		log.Error("CreateEventChain() error = %v, wantErr %v", err)
	}
```
### Create and Sign an event and add it to an existing event chain

```go
    event := &lto.Event{
		Body:      "HeFMDcuveZQYtBePVUugLyWtsiwsW4xp7xKdv",
		Timestamp: time1.Unix(),
		Previous:  "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
		SignKey:   crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
	}
	signedEvent, err := account.SignEvent(event)
	if err != nil {
		log.Error("SignEvent() error = %v", err)
	}
	chain.AddEvent(signedEvent)
```



