package lto_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/ltonetwork/lto-sdk.go/pkg/lto"
)

func Test_NewEventChainCreate(t *testing.T) {
	type fields struct {
		id        []byte
		publicKey []byte
		nonce     []byte
		rand      io.Reader
	}
	type want struct {
		id         []byte
		latestHash string
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name: "should generate a correct hash from the id",
			fields: fields{
				rand: lto.Rand,
				id:   crypto.Base58Decode("L1hGimV7Pp2CFNUnTCitqWDbk9Zng3r3uc66dAG6hLwEx"),
			},
			want: want{
				id:         crypto.Base58Decode("L1hGimV7Pp2CFNUnTCitqWDbk9Zng3r3uc66dAG6hLwEx"),
				latestHash: "9HM1ykH7AxLgdCqBBeUhvoTH4jkq3zsZe4JGTrjXVENg",
			},
			wantErr: false,
		},
		{
			name: "should generate the correct chain id when initiated for an account with random nonce",
			fields: fields{
				rand:      bytes.NewReader(make([]byte, 20)),
				publicKey: crypto.Base58Decode("8MeRTc26xZqPmQ3Q29RJBwtgtXDPwR7P9QNArymjPLVQ"),
			},
			want: want{
				id:         crypto.Base58Decode("2ar3wSjTm1fA33qgckZ5Kxn1x89gKcDPBXTxw56YukdUvrcXXcQh8gKCs8teBh"),
				latestHash: "9y3W6WUsNC72kAa9yeo3kB8b9wULJvBXPRgaHmfXvXjw",
			},
			wantErr: false,
		},
		{
			name: "should generate the correct chain id when initiated for an account with a nonce",
			fields: fields{
				rand:      lto.Rand,
				publicKey: crypto.Base58Decode("8MeRTc26xZqPmQ3Q29RJBwtgtXDPwR7P9QNArymjPLVQ"),
				nonce:     []byte("foo"),
			},
			want: want{
				id:         crypto.Base58Decode("2b6QYLttL2R3CLGL4fUB9vaXXX4c5aFFsoeAmzHWEhqp3bTS49bpomCMTmbV9E"),
				latestHash: "8TJX8LsZCr38uhog9m9YjMF3GNfCDfPCivy6z8Ly5d6f",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldRand := lto.Rand
			lto.Rand = tt.fields.rand
			defer func() { lto.Rand = oldRand }()
			p := lto.NewEventChain()

			if len(tt.fields.id) != 0 {
				p.WithID(tt.fields.id)
			}

			if len(tt.fields.publicKey) != 0 {
				p.WithPublicKey(tt.fields.publicKey)
			}

			if len(tt.fields.nonce) != 0 {
				p.WithNonce(tt.fields.nonce)
			}

			chain, err := p.Create()
			require.NoError(t, err)
			require.Equal(t, tt.want.id, chain.ID)

			hash, err := chain.GetLatestHash()
			require.NoError(t, err)
			require.Equal(t, tt.want.latestHash, hash)
		})
	}
}

func TestEventChain_AddEvent(t *testing.T) {
	type fields struct {
		id           []byte
		body         interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
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
			name: "should add an event and return the latest hash",
			fields: fields{
				id: crypto.Base58Decode("L1hGimV7Pp2CFNUnTCitqWDbk9Zng3r3uc66dAG6hLwEx"),
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			want: &lto.Event{
				Hash: "HxCyzrPmQt7AvWsUoF5JcyKKbcodpUwhsTydvMneAa8h",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event, err := lto.NewEvent().
				WithSignKey(tt.fields.signKey).
				WithTimestamp(tt.fields.timestamp).
				WithBody(tt.fields.body).
				WithPrevious(tt.fields.previousHash).
				Create()
			require.NoError(t, err)

			chain, err := lto.NewEventChain().WithID(tt.fields.id).Create()
			require.NoError(t, err)

			_, err = chain.AddEvent(event)
			require.NoError(t, err)

			hash, err := chain.GetLatestHash()
			require.NoError(t, err)

			require.Equal(t, tt.want.Hash, hash)
		})
	}
}

func TestEventChain_CreateProjectionId(t *testing.T) {
	type fields struct {
		id    []byte
		nonce []byte
		rand  io.Reader
	}
	type args struct {
		event *lto.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "should generate a correct projection id with a set nonce",
			fields: fields{
				id:    crypto.Base58Decode("2b6QYLttL2R3CLGL4fUB9vaXXX4c5HJanjV5QecmAYLCrD52o6is1fRMGShUUF"),
				nonce: []byte("foo"),
			},
			want:    crypto.Base58Decode("2z4AmxL122aaTLyVy6rhEfXHGJMGuXrmahjVCXwYz6GxATR8x3PXNq3XbwbJ2H"),
			wantErr: false,
		},
		{
			name: "should generate a correct projection id with a random nonce",
			fields: fields{
				rand: bytes.NewReader(make([]byte, 20)),
				id:   crypto.Base58Decode("2b6QYLttL2R3CLGL4fUB9vaXXX4c5HJanjV5QecmAYLCrD52o6is1fRMGShUUF"),
			},
			want:    crypto.Base58Decode("2yopB4AaT1phJ4YrXBwbQhimguSM9ZpttRZHMckbf94d3iaERWCPhkAP4quKbs"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldRand := lto.Rand
			lto.Rand = tt.fields.rand
			defer func() { lto.Rand = oldRand }()

			chain, err := lto.NewEventChain().WithID(tt.fields.id).Create()
			require.NoError(t, err)

			projectionID, err := chain.CreateProjectionID(tt.fields.nonce)
			require.NoError(t, err)

			require.Equal(t, tt.want, projectionID)
		})
	}
}
