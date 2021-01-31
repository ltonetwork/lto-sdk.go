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
	log.Error("NewAccount() error = %v", err)
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
	log.Error("NewAccount() error = %v", err)
}
message := []byte("hello")
signedMessage, err := account.SignMessage(message)
if err != nil {
	log.Error("SignMessage() error = %v", err)
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
	log.Error("NewAccount() error = %v", err)
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

## API
### API USAGE
```go
api, err := lto.NewAPI(lto.DefaultMainNetConfig())
if err != nil {
	log.Error("NewAPI() error = %v", err)
}
lto, err := lto.NewClient().WithNetwork(lto.NetworkMain).Create()
if err != nil {
	log.Error("NewClient() error = %v", err)
}
account, err := lto.NewAccount().Create()
if err != nil {
	log.Error("NewAccount() error = %v", err)
}
```
### Balance
#### Fetch Balance

```go
balanceObj, err := api.AddressBalance(account.Address)
if err != nil {
	log.Error("AddressBalance() error = %v", err)
}
fmt.Println("Address", crypto.Base58Encode(balanceObj.Address))
fmt.Println("Balance", balanceObj.Balance)
fmt.Println("Confirmations", balanceObj.Confirmations)
```
#### Balance With Confirmations

```go
balanceObj, err := api.AddressBalanceWithConfirmations(account.Address,10)
if err != nil {
	log.Error("AddressBalanceWithConfirmations() error = %v", err)
}
```
#### Balance Details

```go
balanceDetails, err := api.AddressBalanceDetails(account.Address)
if err != nil {
	log.Error("AddressBalanceDetails() error = %v", err)
}
```
### Blocks

#### Blocks First
```go
blockFirst, err := api.BlocksFirst()
if err != nil {
	log.Error("BlocksFirst() error = %v", err)

}
```
#### Blocks Last

```go
blockLast, err := api.BlocksLast()
if err != nil {
	log.Error("BlocksLast() error = %v", err)

}
```

#### Blocks Height
```go
blockHeight, err := api.BlocksHeight()
if err != nil {
	log.Error("BlocksHeight() error = %v", err)

}
fmt.Println(blockHeight.Height)
```
#### BlocksAt Height 
```go
height:= 100
blocksAt, err := api.BlocksAt(height)
if err != nil {
	log.Error("BlocksAt() error = %v", err)

}
fmt.Println(blocksAt.Signature)
```
#### BlocksGet by Signature
```go
signature := "4DqTwdWnmhgo8Dnwst8px8T37EXubgBKqSHTSGuZ8adDeBVtTS1NBz9PAT3i4HjXsqayT9DWDzcGNoJ5uL5kHKPR"
blocksGet, err := api.BlocksGet(signature)
if err != nil {
	log.Error("BlocksGet() error = %v", err)
}
```

### Transactions
#### Transactions GET
```go
signature := "5C1sBMVCkaS1hr97C5zQCPtvzdF6ubqNauycHesHJ1nyDx7hiaDTdPxwzqJuKNebjho3egWzMVCFxMefNgncSbpp"
transaction, err := api.TransactionsGet(signature)
if err != nil {
	log.Error("TransactionsGet() error = %v", err)

}
```
#### Transactions LIST
```go
address := "3N6mZMgGqYn9EVAR2Vbf637iej4fFipECq8"
limit := 2
transactions, err := api.TransactionsGetList(address, limit)
if err != nil {
	log.Error("TransactionsGetList() error = %v", err)

}
```

#### Transactions UTX Size 
```go
transactionUTX, err := api.TransactionsUTXSize()
if err != nil {
	log.Error("TransactionsUTXSize() error = %v", err)

}
fmt.Println(transactionUTX.Size)
```

#### Transactions UTX GET 
```go
transactionUTX, err := api.TransactionsUTXGet(signature)
if err != nil {
	log.Error("TransactionsUTXSize() error = %v", err)

}
```

#### Transactions UTX LIST 
```go
transactionUTX, err := api.TransactionsUTXGetList()
if err != nil {
	log.Error("TransactionsUTXSize() error = %v", err)

}

```