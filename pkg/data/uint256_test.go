package data

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewUint256FromString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Uint256
		wantErr bool
	}{
		{
			name:  "test 1",
			value: "1000000000000000",
			want: Uint256{
				low:  "0x38d7ea4c68000",
				high: "0x0",
			},
		}, {
			name:  "test 2",
			value: "340282366920938463463374607431768211456",
			want: Uint256{
				low:  "0x0",
				high: "0x1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUint256FromString(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUint256FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint256FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint256_Decimal(t *testing.T) {
	tests := []struct {
		name    string
		uint256 Uint256
		want    decimal.Decimal
		wantErr bool
	}{
		{
			name:    "test 1",
			uint256: NewUint256FromStrings("0x0", "0x1"),
			want:    decimal.RequireFromString("340282366920938463463374607431768211456"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uint256 := Uint256{
				low:  tt.uint256.low,
				high: tt.uint256.high,
			}
			got, err := uint256.Decimal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint256.Decimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint256.Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUint256FromInt(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  Uint256
	}{
		{
			name:  "test 1",
			value: 1,
			want:  NewUint256("0x1", "0x0"),
		}, {
			name:  "test 2",
			value: 257,
			want:  NewUint256("0x101", "0x0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUint256FromInt(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint256FromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
