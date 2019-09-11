package lto_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/ltonetwork/lto-sdk.go/pkg/lto"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/stretchr/testify/require"
)

func TestAccount_GetEncodedPhrase(t *testing.T) {
	type args struct {
		Seed []byte
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
				Seed: []byte("satisfy sustain shiver skill betray mother appear pupil coconut weasel firm top puzzle monkey seek"),
			},
			want:    "EMJxAXyrymyGv1fjRyx9uptWC3Ck5AXxtZbXXv59iDjmV2rQsLmbMmw5DBf1GrjhP9VbE7Dy8wa8VstVnJsXiCDBjJhvUVhyE1wnwA1h9Hdg3wg1V6JFJfszZJ4SxYSuNLQven",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := lto.Account().FromSeed(tt.args.Seed).New()
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
		Sign *crypto.KeyPair
		Seed []byte
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
				Seed: []byte("satisfy sustain shiver skill betray mother appear pupil coconut weasel firm top puzzle monkey seek"),
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
			ap := lto.Account()

			if len(tt.fields.Seed) != 0 {
				ap.FromSeed(tt.fields.Seed)
			} else if tt.fields.Sign != nil {
				ap.FromPrivateKey(tt.fields.Sign.PrivateKey)
			}

			a, err := ap.New()
			require.NoError(t, err)

			got, err := a.SignMessage(tt.args.message)
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
	time1, err := time.Parse(lto.TimeFormat, "2018-03-01T00:00:00+00:00")
	require.NoError(t, err)

	type fields struct {
		Address []byte
		Seed    []byte
		Sign    *crypto.KeyPair
	}
	type args struct {
		event *lto.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *lto.Event
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
				event: &lto.Event{
					Body:      "HeFMDcuveZQYtBePVUugLyWtsiwsW4xp7xKdv",
					Timestamp: time1.Unix(),
					Previous:  "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
					SignKey:   crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
				},
			},
			want: &lto.Event{
				Signature: crypto.Base58Decode("258KnaZxcx4cA9DUWSPw8QwBokRGzFDQmB4BH9MRJhoPJghsXoAZ7KnQ2DWR7ihtjXzUjbsXtSeup4UDcQ2L6RDL"),
				Hash:      "Bpq9rZt12Gv44dkXFw8RmLYzbaH2HBwPQJ6KihdLe5LG",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := lto.Account().FromPrivateKey(tt.fields.Sign.PrivateKey).New()
			require.NoError(t, err)

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
		Sign *crypto.KeyPair
		Seed []byte
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
				Seed: []byte("satisfy sustain shiver skill betray mother appear pupil coconut weasel firm top puzzle monkey seek"),
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
			ap := lto.Account()

			if len(tt.fields.Seed) != 0 {
				ap.FromSeed(tt.fields.Seed)
			} else if tt.fields.Sign != nil {
				ap.FromPrivateKey(tt.fields.Sign.PrivateKey)
			}

			a, err := ap.New()
			require.NoError(t, err)

			got, err := a.Verify(tt.args.signature, tt.args.message)
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

//
//func TestAccount_CreateEventChain(t *testing.T) {
//	type fields struct {
//		Address []byte
//		Seed    []byte
//		Sign    *crypto.KeyPair
//	}
//	type args struct {
//		nonce []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *EventChain
//		wantErr bool
//	}{
//		{
//			name: "should create an account with a random seed",
//			fields: fields{
//				Address: nil,
//				Seed:    nil,
//				Sign: &crypto.KeyPair{
//					PublicKey:  nil,
//					PrivateKey: nil,
//				},
//			},
//			args: args{
//				nonce: nil,
//			},
//			want: &EventChain{
//				ID:     nil,
//				Events: nil,
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			account, err := Account().FromSeed(tt.fields.Seed).New()
//			require.NoError(t, err)
//
//			got, err := account.CreateEventChain(tt.args.nonce)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CreateEventChain() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreateEventChain() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
