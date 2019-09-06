package api

import (
	"reflect"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/stretchr/testify/require"
)

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
			name: "should generate a correct event chain id with a nonce given",
			fields: fields{
				NetworkByte: MainNetByte,
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
			lto, err := NewLTO(tt.fields.NetworkByte, "")
			require.NoError(t, err)

			got, err := lto.CreateEventChainID(tt.args.publicSignKey, tt.args.nonce)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEventChainID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEventChainID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
