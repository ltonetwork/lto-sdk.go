package api

import (
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/require"
)

func TestAPI_Balance(t *testing.T) {
	type fields struct {
		config *LTOConfig
	}
	type args struct {
		confirmations int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *BalanceResponse
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				config: DefaultTestNetConfig,
			},
			want: &BalanceResponse{
				Confirmations: 0,
				Balance:       0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := NewAPI(tt.fields.config)
			require.NoError(t, err)

			lto, err := NewLTO(TestNetByte, "")
			require.NoError(t, err)

			a, err := lto.CreateAccount(15)
			require.NoError(t, err)

			spew.Dump(a.Address)

			got, err := api.Balance(a.Address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			spew.Dump(crypto.Base58Encode(got.Address))
			spew.Dump(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Balance() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestAPI_BalanceWithConfirmations(t *testing.T) {
	type fields struct {
		config *LTOConfig
	}
	type args struct {
		confirmations int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *BalanceResponse
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				config: DefaultTestNetConfig,
			},
			want: &BalanceResponse{
				Confirmations: 0,
				Balance:       0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := NewAPI(tt.fields.config)
			require.NoError(t, err)

			lto, err := NewLTO(TestNetByte, "")
			require.NoError(t, err)

			a, err := lto.CreateAccount(15)
			require.NoError(t, err)

			spew.Dump(a.Address)

			got, err := api.BalanceWithConfirmations(a.Address, 10)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			spew.Dump(crypto.Base58Encode(got.Address))
			spew.Dump(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Balance() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestAPI_BalanceDetails(t *testing.T) {
	type fields struct {
		config *LTOConfig
	}
	type args struct {
		confirmations int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *BalanceResponse
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				config: DefaultTestNetConfig,
			},
			want: &BalanceResponse{
				Confirmations: 0,
				Balance:       0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := NewAPI(tt.fields.config)
			require.NoError(t, err)

			lto, err := NewLTO(TestNetByte, "")
			require.NoError(t, err)

			a, err := lto.CreateAccount(15)
			require.NoError(t, err)

			spew.Dump(a.Address)

			got, err := api.BalanceDetails(a.Address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			spew.Dump(crypto.Base58Encode(got.Address))
			spew.Dump(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Balance() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
