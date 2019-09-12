package lto

import (
	"testing"

	"github.com/ltonetwork/lto-sdk.go/pkg/crypto"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/require"
)

func TestAPI_Balance(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			lto, err := NewClient().WithNetwork(NetworkTest).Create()
			require.NoError(t, err)

			a, err := lto.NewAccount().Create()
			require.NoError(t, err)

			spew.Dump(a.Address)

			got, err := api.AddressBalance(a.Address)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			spew.Dump(crypto.Base58Encode(got.Address))
			spew.Dump(got)
		})
	}
}

func TestAPI_BalanceWithConfirmations(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			lto, err := NewClient().WithNetwork(NetworkTest).Create()
			require.NoError(t, err)

			a, err := lto.NewAccount().Create()
			require.NoError(t, err)

			spew.Dump(a.Address)

			got, err := api.AddressBalanceWithConfirmations(a.Address, 10)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			spew.Dump(crypto.Base58Encode(got.Address))
			spew.Dump(got)
		})
	}
}

func TestAPI_BalanceDetails(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			lto, err := NewClient().WithNetwork(NetworkTest).Create()
			require.NoError(t, err)

			a, err := lto.NewAccount().Create()
			require.NoError(t, err)

			spew.Dump(a.Address)

			got, err := api.AddressBalanceDetails(a.Address)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			spew.Dump(crypto.Base58Encode(got.Address))
			spew.Dump(got)
		})
	}
}

func TestAPI_BlocksLast(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.BlocksLast()
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_BlocksFirst(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.BlocksFirst()
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_BlocksHeight(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.BlocksHeight()
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_BlocksAt(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.BlocksAt(100)
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_BlocksGet(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.BlocksGet("hXTXbA4wWtM7HSZqd4YZbYAFfVqTcAtcXzjEVAwxcJoGuC46C8t3ewiKdUybHsdBk6YsPUUzY3Zkr7Ww9qejzPh")
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_UtilsTime(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.UtilsTime()
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_UtilsCompile(t *testing.T) {
	t.Skip()
	// TODO Find an example for a valid script
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.UtilsCompile("")
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_TransactionsGet(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.TransactionsGet("5C1sBMVCkaS1hr97C5zQCPtvzdF6ubqNauycHesHJ1nyDx7hiaDTdPxwzqJuKNebjho3egWzMVCFxMefNgncSbpp")
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_TransactionsGetList(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.TransactionsGetList("3N6mZMgGqYn9EVAR2Vbf637iej4fFipECq8", 2)
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_TransactionsUTXSize(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.TransactionsUTXSize()
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}

func TestAPI_TransactionsUTXGet(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.TransactionsUTXGet("asd")
			require.Error(t, err)
			spew.Dump(res)
		})
	}
}

func TestAPI_TransactionsUTXGetList(t *testing.T) {
	type fields struct {
		config *Config
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
				config: DefaultTestNetConfig(),
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

			res, err := api.TransactionsUTXGetList()
			require.NoError(t, err)
			require.NotNil(t, res)
			spew.Dump(res)
		})
	}
}
