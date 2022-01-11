package unitconv

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/synapsecns/synapsecommon-go/bigint"
)

func TestParseEther(t *testing.T) {
	weiBigInt, ok := new(big.Int).SetString("744000000000000000000", 10)
	assert.Truef(t, ok, "error creating big.Int from string %[1]s", "744000000000000000000")

	type args struct {
		amt      interface{}
		fromUnit Unit
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Float
		wantErr bool
	}{
		{
			"from gwei string",
			args{"744", GWei},
			newEther().SetFloat64(0.000000744),
			false,
		},
		{
			"from gwei float64",
			args{float64(744), GWei},
			newEther().SetFloat64(0.000000744),
			false,
		},
		{
			"from gwei big.Int",
			args{bigint.FromInt64(744), GWei},
			newEther().SetFloat64(0.000000744),
			false,
		},
		{
			"from gwei big.Float",
			args{
				new(big.Float).SetInt64(744).SetPrec(EtherDecimalPrecision),
				GWei,
			},
			newEther().SetFloat64(0.000000744),
			false,
		},
		{
			"from wei string",
			args{"744000000000000000000", Wei},
			newEther().SetFloat64(744),
			false,
		},
		{
			"from wei big.Int",
			args{weiBigInt, Wei},
			newEther().SetFloat64(744),
			false,
		},
		{
			"from szabo big.Int",
			args{big.NewInt(744000000), Szabo},
			newEther().SetFloat64(744),
			false,
		},
		{
			"from szabo uint64",
			args{uint64(744000), Szabo},
			newEther().SetFloat64(.744),
			false,
		},
		{
			"from TEther",
			args{"0.000000000744", TEther},
			newEther().SetFloat64(744),
			false,
		},
		{
			"invalid type",
			args{true, KWei},
			nil,
			true,
		},
		{
			"invalid Unit",
			args{1, Unit(744)},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToEther(tt.args.amt, tt.args.fromUnit)

			if tt.wantErr {
				if !assert.Errorf(t, err, "ParseEther() error = %v, wantErr %v", err, tt.wantErr) {
					t.FailNow()
				} else {
					return
				}
			} else {
				if !assert.NoErrorf(t, err, "ParseEther() error = %v, wantErr %v", err, tt.wantErr) {
					t.FailNow()
				}
			}

			assert.Equalf(t, tt.want.Cmp(got), 0, "ParseEther() got = %v, want %v", got, tt.want)
		})
	}
}

func TestParseWei(t *testing.T) {
	weiBigInt, ok := new(big.Int).SetString("744000000000000000000", 10)
	if !assert.Truef(t, ok, "error creating big.Int from string %[1]s", "744000000000000000000") {
		t.FailNow()
	}

	type args struct {
		amt      interface{}
		fromUnit Unit
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			"from gwei string",
			args{"744", GWei},
			bigint.FromUint64(744000000000),
			false,
		},
		{
			"from gwei float64",
			args{float64(744), GWei},
			bigint.FromUint64(744000000000),
			false,
		},
		{
			"from gwei big.Int",
			args{bigint.FromInt64(744), GWei},
			bigint.FromUint64(744000000000),
			false,
		},
		{
			"from gwei big.Float",
			args{
				new(big.Float).SetInt64(744).SetPrec(EtherDecimalPrecision),
				GWei,
			},
			bigint.FromUint64(744000000000),
			false,
		},
		{
			"from wei string",
			args{"744000000000000000000", Wei},
			weiBigInt,
			false,
		},
		{
			"from wei big.Int",
			args{weiBigInt, Wei},
			weiBigInt,
			false,
		},
		{
			"from szabo big.Int",
			args{big.NewInt(744000000), Szabo},
			weiBigInt,
			false,
		},
		{
			"from szabo string",
			args{"744000000", Szabo},
			weiBigInt,
			false,
		},
		{
			"from TEther",
			args{"0.000000000744", TEther},
			weiBigInt,
			false,
		},
		{
			"from Ether string",
			args{"744", Ether},
			weiBigInt,
			false,
		},
		{
			"from Ether big.Float",
			args{big.NewFloat(744), Ether},
			weiBigInt,
			false,
		},
		{
			"invalid type",
			args{true, KWei},
			nil,
			true,
		},
		{
			"invalid Unit",
			args{1, Unit(744)},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToWei(tt.args.amt, tt.args.fromUnit)

			if tt.wantErr {
				if !assert.Errorf(t, err, "ParseWei(%v, %v)", tt.args.amt, tt.args.fromUnit) {
					t.FailNow()
				} else {
					return
				}
			} else {
				if !assert.NoErrorf(t, err, "ParseWei(%v, %v)", tt.args.amt, tt.args.fromUnit) {
					t.FailNow()
				}
			}

			assert.NotNil(t, got)

			assert.Equalf(t, tt.want.Cmp(got), 0, "ParseWei() got = %v, want %v", got, tt.want)
		})
	}
}
