package data

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestDAMode_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		want    DAMode
		text    []byte
		wantErr bool
	}{
		{
			name: "test 1",
			want: DAModeL1,
			text: []byte(`{"mode":0}`),
		}, {
			name: "test 2",
			want: DAModeL2,
			text: []byte(`{"mode":1}`),
		}, {
			name: "test 3",
			want: DAModeL1,
			text: []byte(`{"mode":"L1"}`),
		}, {
			name: "test 4",
			want: DAModeL2,
			text: []byte(`{"mode":"L2"}`),
		}, {
			name:    "test 5",
			text:    []byte(`{"mode":10}`),
			wantErr: true,
		}, {
			name:    "test 6",
			text:    []byte(`{"mode":"10"}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type buf struct {
				Mode DAMode `json:"mode"`
			}
			var x buf
			err := json.Unmarshal(tt.text, &x)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, x.Mode)
			}
		})
	}
}
