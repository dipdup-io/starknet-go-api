package data

import (
	"testing"
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
		name    string
		f       Felt
		want    string
		wantErr bool
	}{
		{
			name: "test 1",
			f:    Felt("0x7572692f706963742f7433382e6a7067"),
			want: "uri/pict/t38.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.ToAsciiString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Felt.ToAsciiString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Felt.ToAsciiString() = %v, want %v", got, tt.want)
			}
		})
	}
}
