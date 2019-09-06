package api

import (
	"bytes"
	cryptorand "crypto/rand"
	"io"
	"reflect"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/stretchr/testify/require"
)

func TestLTO_CreateAccount(t *testing.T) {
	type fields struct {
		NetworkByte byte
	}
	type args struct {
		words int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantLenPrivate int
		wantLenPublic  int
		wantErr        bool
	}{
		{
			name: "should return true for a valid address",
			fields: fields{
				NetworkByte: MainNetByte,
			},
			args: args{
				words: 0,
			},
			wantLenPrivate: 64,
			wantLenPublic:  32,
		},
		{
			name: "should return false for an invalid address",
			fields: fields{
				NetworkByte: MainNetByte,
			},
			args: args{
				words: 15,
			},
			wantLenPrivate: 64,
			wantLenPublic:  32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lto, err := NewLTO(tt.fields.NetworkByte, "")
			require.NoError(t, err)

			got, err := lto.CreateAccount(tt.args.words)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				require.Len(t, got.Sign.PrivateKey, tt.wantLenPrivate)
				require.Len(t, got.Sign.PublicKey, tt.wantLenPublic)
			}
		})
	}
}

func TestLTO_CreateAccountFromExistingPhrase(t *testing.T) {
	type fields struct {
		NetworkByte byte
	}
	type args struct {
		phrase []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "should create an account with from an existing seed",
			fields: fields{
				NetworkByte: MainNetByte,
			},
			args: args{
				phrase: []byte("manage manual recall harvest series desert melt police rose hollow moral pledge kitten position add"),
			},
			want: &Account{
				Address: crypto.Base58Decode("3JmCa4jLVv7Yn2XkCnBUGsa7WNFVEMxAfWe"),
				Sign: &crypto.KeyPair{
					PublicKey:  crypto.Base58Decode("GjSacB6a5DFNEHjDSmn724QsrRStKYzkahPH67wyrhAY"),
					PrivateKey: crypto.Base58Decode("4zsR9xoFpxfnNwLcY4hdRUarwf5xWtLj6FpKGDFBgscPxecPj2qgRNx4kJsFCpe9YDxBRNoeBWTh2SDAdwTySomS"),
				},
			},
			wantErr: false,
		},
		{
			name: "should create an account with from an existing seed for testnet",
			fields: fields{
				NetworkByte: TestNetByte,
			},
			args: args{
				phrase: []byte("manage manual recall harvest series desert melt police rose hollow moral pledge kitten position add"),
			},
			want: &Account{
				Address: crypto.Base58Decode("3MyuPwbiobZFnZzrtyY8pkaHoQHYmyQxxY1"),
				Sign: &crypto.KeyPair{
					PublicKey:  crypto.Base58Decode("GjSacB6a5DFNEHjDSmn724QsrRStKYzkahPH67wyrhAY"),
					PrivateKey: crypto.Base58Decode("4zsR9xoFpxfnNwLcY4hdRUarwf5xWtLj6FpKGDFBgscPxecPj2qgRNx4kJsFCpe9YDxBRNoeBWTh2SDAdwTySomS"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lto, err := NewLTO(tt.fields.NetworkByte, "")
			require.NoError(t, err)

			got, err := lto.CreateAccountFromExistingPhrase(tt.args.phrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccountFromExistingPhrase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				require.Equal(t, tt.want.Address, got.Address)
				require.Equal(t, tt.want.Sign, got.Sign)
			}
		})
	}
}

func TestLTO_CreateAccountFromPrivateKey(t *testing.T) {
	type fields struct {
		NetworkByte byte
	}
	type args struct {
		privateKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "should create an account with from an existing private key",
			fields: fields{
				NetworkByte: MainNetByte,
			},
			args: args{
				privateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
			},
			want: &Account{
				Sign: &crypto.KeyPair{
					PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
					PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lto, err := NewLTO(tt.fields.NetworkByte, "")
			require.NoError(t, err)

			got, err := lto.CreateAccountFromPrivateKey(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccountFromExistingPhrase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				require.Equal(t, tt.want.Sign, got.Sign)
			}
		})
	}
}

func TestLTO_CreateEventChainID(t *testing.T) {
	type fields struct {
		NetworkByte byte
		Random      io.Reader
	}
	type args struct {
		publicSignKey []byte
		nonce         []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "should generate a correct event chain id without a nonce",
			fields: fields{
				NetworkByte: MainNetByte,
				Random:      bytes.NewReader(make([]byte, 20)),
			},
			args: args{
				publicSignKey: crypto.Base58Decode("8MeRTc26xZqPmQ3Q29RJBwtgtXDPwR7P9QNArymjPLVQ"),
				nonce:         nil,
			},
			want:    crypto.Base58Decode("2ar3wSjTm1fA33qgckZ5Kxn1x89gKcDPBXTxw56YukdUvrcXXcQh8gKCs8teBh"),
			wantErr: false,
		},
		{
			name: "should generate a correct event chain id with a nonce given",
			fields: fields{
				NetworkByte: MainNetByte,
				Random:      cryptorand.Reader,
			},
			args: args{
				publicSignKey: crypto.Base58Decode("8MeRTc26xZqPmQ3Q29RJBwtgtXDPwR7P9QNArymjPLVQ"),
				nonce:         []byte("foo"),
			},
			want:    crypto.Base58Decode("2b6QYLttL2R3CLGL4fUB9vaXXX4c5aFFsoeAmzHWEhqp3bTS49bpomCMTmbV9E"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldRand := rand
			rand = tt.fields.Random

			lto, err := NewLTO(tt.fields.NetworkByte, "")
			require.NoError(t, err)

			got, err := lto.CreateEventChainID(tt.args.publicSignKey, tt.args.nonce)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEventChainID() error = %v, wantErr %v", err, tt.wantErr)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEventChainID() got = %v, want %v", got, tt.want)
			}

			rand = oldRand
		})
	}
}

func TestLTO_IsValidAddress(t *testing.T) {
	type fields struct {
		NetworkByte byte
	}
	type args struct {
		address []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should return true for a valid address",
			fields: fields{
				NetworkByte: MainNetByte,
			},
			args: args{
				address: crypto.Base58Decode("3JmCa4jLVv7Yn2XkCnBUGsa7WNFVEMxAfWe"),
			},
			want: true,
		},
		{
			name: "should return false for an invalid address",
			fields: fields{
				NetworkByte: MainNetByte,
			},
			args: args{
				address: crypto.Base58Decode("3JmCa4jLVv7Yn2XkCnBUGsa7WNFVEMxAfW1"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lto, err := NewLTO(tt.fields.NetworkByte, "")
			require.NoError(t, err)

			if got := lto.IsValidAddress(tt.args.address); got != tt.want {
				t.Errorf("IsValidAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
