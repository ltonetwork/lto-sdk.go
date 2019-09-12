package lto_test

import (
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/stretchr/testify/require"

	"github.com/ltonetwork/lto-sdk.go/pkg/lto"
)

type Data struct {
	Foo   string `json:"foo"`
	Color string `json:"color"`
}

func Test_NewEventCreate(t *testing.T) {
	type fields struct {
		body         interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    *lto.Event
		wantErr bool
	}{
		{
			name: "should create a correct event object",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			want: &lto.Event{
				Body:      "HeFMDcuveZQYtBePVUugLyWtsiwsW4xp7xKdv",
				Timestamp: 1519862400,
				Previous:  "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
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

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, tt.want.Body, event.Body)
			require.Equal(t, tt.want.Timestamp, event.Timestamp)
			require.Equal(t, tt.want.Previous, event.Previous)
		})
	}
}

func TestEvent_GetMessage(t *testing.T) {
	type fields struct {
		body         interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "should create a correct event object",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			want: []byte(`HeFMDcuveZQYtBePVUugLyWtsiwsW4xp7xKdv
1519862400
72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW
FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y`),
			wantErr: false,
		},
		{
			name: "should throw an error when no body is set",
			fields: fields{
				body:         nil,
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			wantErr: true,
		},
		{
			name: "should throw an error when no signkey is set",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
			},
			wantErr: true,
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
			message, err := event.GetMessage()

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, string(tt.want), string(message))
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestEvent_GetHash(t *testing.T) {
	type fields struct {
		body         interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should generate a correct hash",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			want:    "Bpq9rZt12Gv44dkXFw8RmLYzbaH2HBwPQJ6KihdLe5LG",
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
			message, err := event.GetHash()

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, message)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestEvent_GetBody(t *testing.T) {
	type fields struct {
		body         interface{}
		bodyArg      interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "should return a decoded body",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				bodyArg:      new(Data),
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			want: &Data{
				Foo:   "bar",
				Color: "red",
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

			err = event.GetBody(tt.fields.bodyArg)

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, tt.fields.bodyArg)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestEvent_VerifySignature(t *testing.T) {
	type fields struct {
		body         interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
		signature    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "should verify a correctly signed event",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
				signature:    crypto.Base58Decode("258KnaZxcx4cA9DUWSPw8QwBokRGzFDQmB4BH9MRJhoPJghsXoAZ7KnQ2DWR7ihtjXzUjbsXtSeup4UDcQ2L6RDL"),
			},
			want:    true,
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
				WithSignature(tt.fields.signature).
				Create()
			require.NoError(t, err)
			message, err := event.VerifySignature()

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, message)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestEvent_SignWith(t *testing.T) {
	type fields struct {
		body         interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
		signature    []byte
		privateKey   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    *lto.Event
		wantErr bool
	}{
		{
			name: "should generate a correct signature",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				privateKey:   crypto.Base58Decode("wJ4WH8dD88fSkNdFQRjaAhjFUZzZhV5yiDLDwNUnp6bYwRXrvWV8MJhQ9HL9uqMDG1n7XpTGZx7PafqaayQV8Rp"),
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
				signature:    crypto.Base58Decode("258KnaZxcx4cA9DUWSPw8QwBokRGzFDQmB4BH9MRJhoPJghsXoAZ7KnQ2DWR7ihtjXzUjbsXtSeup4UDcQ2L6RDL"),
			},
			want: &lto.Event{
				SignKey:   crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
				Signature: crypto.Base58Decode("258KnaZxcx4cA9DUWSPw8QwBokRGzFDQmB4BH9MRJhoPJghsXoAZ7KnQ2DWR7ihtjXzUjbsXtSeup4UDcQ2L6RDL"),
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
				WithSignature(tt.fields.signature).
				Create()
			require.NoError(t, err)

			a, err := lto.NewAccount().FromPrivateKey(tt.fields.privateKey).Create()
			require.NoError(t, err)

			event, err = event.SignWith(a)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want.Signature, event.Signature)
				require.Equal(t, tt.want.SignKey, event.SignKey)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestEvent_AddTo(t *testing.T) {
	type fields struct {
		body         interface{}
		bodyArg      interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "should return a decoded body",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				bodyArg:      new(Data),
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			want: &Data{
				Foo:   "bar",
				Color: "red",
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

			chain, err := lto.NewEventChain().Create()
			require.NoError(t, err)

			event, err = event.AddTo(chain)
			require.NoError(t, err)
			require.Contains(t, chain.Events, event)
		})
	}
}

func TestEvent_GetResourceVersion(t *testing.T) {
	type fields struct {
		body         interface{}
		bodyArg      interface{}
		previousHash string
		timestamp    int64
		signKey      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "should return a decoded body",
			fields: fields{
				body: &Data{
					Foo:   "bar",
					Color: "red",
				},
				bodyArg:      new(Data),
				timestamp:    1519862400,
				previousHash: "72gRWx4C1Egqz9xvUBCYVdgh7uLc5kmGbjXFhiknNCTW",
				signKey:      crypto.Base58Decode("FkU1XyfrCftc4pQKXCrrDyRLSnifX1SMvmx1CYiiyB3Y"),
			},
			want:    "4RaPGFmq",
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

			v := event.GetResourceVersion()
			require.Equal(t, tt.want, v)
		})
	}
}
