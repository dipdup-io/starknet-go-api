package data

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func Test_NewFromAsciiString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Felt
	}{
		{
			name: "test 1",
			s:    "uri/pict/t38.jpg",
			want: Felt("0x7572692f706963742f7433382e6a7067"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromAsciiString(tt.s); got != tt.want {
				t.Errorf("Felt.ToShortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFelt_ToAsciiString(t *testing.T) {
	tests := []struct {
		name string
		f    Felt
		want string
	}{
		{
			name: "test 1",
			f:    Felt("0x7572692f706963742f7433382e6a7067"),
			want: "uri/pict/t38.jpg",
		}, {
			name: "test 2",
			f:    Felt("0x1"),
			want: "",
		}, {
			name: "test 3",
			f:    Felt("0x0"),
			want: string([]byte{0x00}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.f.ToAsciiString()
			if got != tt.want {
				t.Errorf("Felt.ToAsciiString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFelt_Decimal(t *testing.T) {
	tests := []struct {
		name string
		f    Felt
		want decimal.Decimal
	}{
		{
			name: "test 1",
			f:    Felt("0x0362a8f174e36882a3d8da3ba18e85ff9da5aec5019c8782e3a9273ca7858aa3"),
			want: decimal.RequireFromString("1531255561141264586217574535501033622415882562452753316784814711352206920355"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Decimal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Felt.Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
