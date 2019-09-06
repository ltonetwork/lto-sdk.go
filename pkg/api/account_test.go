package api

import (
	"reflect"
	"testing"
	"time"

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
		phrase      []byte
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
				phrase:      []byte("satisfy sustain shiver skill betray mother appear pupil coconut weasel firm top puzzle monkey seek"),
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

			if len(tt.fields.phrase) != 0 {
				account, err = NewAccount(tt.fields.phrase, tt.fields.NetworkByte)
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

func TestAccount_SignEvent(t *testing.T) {
	time1, err := time.Parse(TimeFormat, "2018-03-01T00:00:00+00:00")
	require.NoError(t, err)

	type fields struct {
		Address []byte
		Seed    []byte
		Sign    *crypto.KeyPair
	}
	type args struct {
		event *Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name: "should create a correct signature for an event",
			fields: fields{
				Sign: &crypto.KeyPair{
					PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
					PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
				},
			},
			args: args{
				event: &Event{
					Body:      "HeFMDcuveZQYtBePVUugLyWtsiwsW4xp7xKdv",
					Timestamp: time1.Unix(),
					Previous:  "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
					SignKey:   crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
				},
			},
			want: &Event{
				Signature: crypto.Base58Decode("258KnaZxcx4cA9DUWSPw8QwBokRGzFDQmB4BH9MRJhoPJghsXoAZ7KnQ2DWR7ihtjXzUjbsXtSeup4UDcQ2L6RDL"),
				Hash:      "Bpq9rZt12Gv44dkXFw8RmLYzbaH2HBwPQJ6KihdLe5LG",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Sign: tt.fields.Sign,
			}
			got, err := a.SignEvent(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				require.Equal(t, tt.want.Hash, got.Hash)
				require.Equal(t, tt.want.Signature, got.Signature)
			}
		})
	}
}

func TestAccount_Verify(t *testing.T) {
	type fields struct {
		Sign        *crypto.KeyPair
		phrase      []byte
		networkByte byte
	}
	type args struct {
		signature []byte
		message   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "should verify a correct signature to be true",
			fields: fields{
				Sign: &crypto.KeyPair{
					PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
					PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
				},
			},
			args: args{
				signature: crypto.Base58Decode("2DDGtVHrX66Ae8C4shFho4AqgojCBTcE4phbCRTm3qXCKPZZ7reJBXiiwxweQAkJ3Tsz6Xd3r5qgnbA67gdL5fWE"),
				message:   []byte("hello"),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "should verify a correct signature with seeded account to be true",
			fields: fields{
				phrase:      []byte("satisfy sustain shiver skill betray mother appear pupil coconut weasel firm top puzzle monkey seek"),
				networkByte: 0,
			},
			args: args{
				signature: crypto.Base58Decode("2SPPcJzvJHTNJWjzWLWDaaiZap61L5EwhPY9fRjLTqGebDuqoCuqGCVTTQVyAiMAeffuNXbR8oBNRdauSr63quhn"),
				message:   []byte("hello"),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "should verify an incorrect signature to be false",
			fields: fields{
				Sign: &crypto.KeyPair{
					PublicKey:  crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
					PrivateKey: crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
				},
			},
			args: args{
				signature: crypto.Base58Decode("2DDGtVHrX66Ae8C4shFho4AqgojCBTcE4phbCRTm3qXCKPZZ7reJBXiiwxweQAkJ3Tsz6Xd3r5qgnbA67gdL5fWE"),
				message:   []byte("not this"),
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			var account *Account

			if len(tt.fields.phrase) != 0 {
				account, err = NewAccount(tt.fields.phrase, tt.fields.networkByte)
				require.NoError(t, err)
			} else {
				account = &Account{
					Sign: tt.fields.Sign,
				}
			}

			got, err := account.Verify(tt.args.signature, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}
