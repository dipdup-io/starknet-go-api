package api

import (
	"reflect"
	"testing"
)

func TestSyncing_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Syncing
		wantErr bool
	}{
		{
			name: "state",
			data: []byte(`{
				"starting_block_num": "0x9bf0c",
				"current_block_num": "0xb438e",
				"highest_block_num": "0xb438e",
				"starting_block_hash": "0x368c35989aa3b8e43995f5820bc5eff1ae2411113cbc38244ca80c372026b99",
				"current_block_hash": "0x6b391105913b7093bb26109ad1a6c29f7ab4302f6b8202ef142608529686d4a",
				"highest_block_hash": "0x6b391105913b7093bb26109ad1a6c29f7ab4302f6b8202ef142608529686d4a"
			  }`),
			want: Syncing{
				StartingBlockNum:  "0x9bf0c",
				CurrentBlockNum:   "0xb438e",
				HighestBlockNum:   "0xb438e",
				StartingBlockHash: "0x368c35989aa3b8e43995f5820bc5eff1ae2411113cbc38244ca80c372026b99",
				CurrentBlockHash:  "0x6b391105913b7093bb26109ad1a6c29f7ab4302f6b8202ef142608529686d4a",
				HighestBlockHash:  "0x6b391105913b7093bb26109ad1a6c29f7ab4302f6b8202ef142608529686d4a",
				Synced:            true,
			},
		}, {
			name: "bool",
			data: []byte(`false`),
			want: Syncing{
				Synced: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Syncing
			if err := s.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("Syncing.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.want, s) {
				t.Errorf("Syncing.UnmarshalJSON() got=%v want=%v", s, tt.want)
			}
		})
	}
}
