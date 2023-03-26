package encoding

import "testing"

func TestGetSelectorFromName(t *testing.T) {
	tests := []struct {
		name       string
		entrypoint string
		want       string
	}{
		{
			name:       "constructor",
			entrypoint: "constructor",
			want:       "28ffe4ff0f226a9107253e17a904099aa4f63a02a5621de0576e5aa71bc5194",
		}, {
			name:       "__execute__",
			entrypoint: "__execute__",
			want:       "15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad",
		}, {
			name:       "__validate__",
			entrypoint: "__validate__",
			want:       "162da33a4585851fe8d3af3c2a9c60b557814e221e0d4f30ff0b2189d9c7775",
		}, {
			name:       "__validate_declare__",
			entrypoint: "__validate_declare__",
			want:       "289da278a8dc833409cabfdad1581e8e7d40e42dcaed693fa4008dcdb4963b3",
		}, {
			name:       "__validate_deploy__",
			entrypoint: "__validate_deploy__",
			want:       "36fcbf06cd96843058359e1a75928beacfac10727dab22a3972f0af8aa92895",
		}, {
			name:       "transfer",
			entrypoint: "transfer",
			want:       "83afd3f4caedc6eebf44246fe54e38c95e3179a5ec9ea81740eca5b482d12e",
		}, {
			name:       "call_xor_counters",
			entrypoint: "call_xor_counters",
			want:       "30f842021fbf02caf80d09a113997c1e00a32870eee0c6136bed27acb348bea",
		}, {
			name:       "__l1_default__",
			entrypoint: "__l1_default__",
			want:       "1ac445d7a589d19896ced0342d3d23fc7a05f2d7feaf2cd89e84104c86b5937",
		}, {
			name:       "mint",
			entrypoint: "mint",
			want:       "2f0b3c5710379609eb5495f1ecd348cb28167711b73609fe565a72734550354",
		}, {
			name:       "__default__",
			entrypoint: "__default__",
			want:       "2e4c01ac72b840834c6c3146a782496a90a442ac831e5188090c1d33a7c0f50",
		}, {
			name:       "set_implementation",
			entrypoint: "set_implementation",
			want:       "c4e105e5276c704b5490fa2ab565b6b1904912203fbc6e7bcdeb51fa8c1ef2",
		}, {
			name:       "Transfer",
			entrypoint: "Transfer",
			want:       "99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9",
		}, {
			name:       "balanceOf",
			entrypoint: "balanceOf",
			want:       "2e4263afad30923c891518314c3c95dbe830a16874e8abc5777a9a20b54c76e",
		}, {
			name:       "Approval",
			entrypoint: "Approval",
			want:       "134692b230b9e1ffa39098904722134159652b09c5bc41d88d6698779d228ff",
		}, {
			name:       "changeModules",
			entrypoint: "changeModules",
			want:       "3ffada7235f48d4811be030385f19e6d50e2cfa368ded42f1892666f834e407",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSelectorFromName(tt.entrypoint); got != tt.want {
				t.Errorf("GetSelectorFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}
