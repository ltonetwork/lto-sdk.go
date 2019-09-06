package api

import (
	"reflect"
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/stretchr/testify/require"
)

func TestAccount_GetEncodedPhrase(t *testing.T) {
	type args struct {
		phrase      []byte
		networkByte byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return a correct base58 encoded phrase",
			args: args{
				phrase:      []byte("satisfy sustain shiver skill betray mother appear pupil coconut weasel firm top puzzle monkey seek"),
				networkByte: 0,
			},
			want:    "EMJxAXyrymyGv1fjRyx9uptWC3Ck5AXxtZbXXv59iDjmV2rQsLmbMmw5DBf1GrjhP9VbE7Dy8wa8VstVnJsXiCDBjJhvUVhyE1wnwA1h9Hdg3wg1V6JFJfszZJ4SxYSuNLQven",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := NewAccount(tt.args.phrase, tt.args.networkByte)
			require.NoError(t, err)

			got := account.GetEncodedPhrase()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEncodedPhrase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_SignMessage(t *testing.T) {
	type fields struct {
		Sign        *crypto.KeyPair
		Phrase      []byte
		NetworkByte byte
	}
	type args struct {
		message []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "should generate a correct signature from a message",
			fields: fields{
				Sign: &crypto.KeyPair{
					PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
					PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
				},
			},
			args: args{
				message: []byte("hello"),
			},
			want:    crypto.Base58Decode("2DDGtVHrX66Ae8C4shFho4AqgojCBTcE4phbCRTm3qXCKPZZ7reJBXiiwxweQAkJ3Tsz6Xd3r5qgnbA67gdL5fWE"),
			wantErr: false,
		},
		{
			name: "should generate a correct signature from a message with a seeded account",
			fields: fields{
				Phrase:      []byte("satisfy sustain shiver skill betray mother appear pupil coconut weasel firm top puzzle monkey seek"),
				NetworkByte: 0,
			},
			args: args{
				message: []byte("hello"),
			},
			want:    crypto.Base58Decode("2SPPcJzvJHTNJWjzWLWDaaiZap61L5EwhPY9fRjLTqGebDuqoCuqGCVTTQVyAiMAeffuNXbR8oBNRdauSr63quhn"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			var account *Account

			if len(tt.fields.Phrase) != 0 {
				account, err = NewAccount(tt.fields.Phrase, tt.fields.NetworkByte)
				require.NoError(t, err)
			} else {
				account = &Account{
					Sign: tt.fields.Sign,
				}
			}

			got, err := account.SignMessage(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
