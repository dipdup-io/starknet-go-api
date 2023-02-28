package encoding

import "testing"

func TestGetSelectorFromName(t *testing.T) {
	tests := []struct {
		name     string
		selector string
		want     string
	}{
		{
			name:     "constructor",
			selector: "constructor",
			want:     "28ffe4ff0f226a9107253e17a904099aa4f63a02a5621de0576e5aa71bc5194",
		}, {
			name:     "__execute__",
			selector: "__execute__",
			want:     "15d40a3d6ca2ac30f4031e42be28da9b056fef9bb7357ac5e85627ee876e5ad",
		}, {
			name:     "__validate__",
			selector: "__validate__",
			want:     "162da33a4585851fe8d3af3c2a9c60b557814e221e0d4f30ff0b2189d9c7775",
		}, {
			name:     "__validate_declare__",
			selector: "__validate_declare__",
			want:     "289da278a8dc833409cabfdad1581e8e7d40e42dcaed693fa4008dcdb4963b3",
		}, {
			name:     "__validate_deploy__",
			selector: "__validate_deploy__",
			want:     "36fcbf06cd96843058359e1a75928beacfac10727dab22a3972f0af8aa92895",
		}, {
			name:     "transfer",
			selector: "transfer",
			want:     "83afd3f4caedc6eebf44246fe54e38c95e3179a5ec9ea81740eca5b482d12e",
		}, {
			name:     "call_xor_counters",
			selector: "call_xor_counters",
			want:     "30f842021fbf02caf80d09a113997c1e00a32870eee0c6136bed27acb348bea",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSelectorFromName(tt.selector); got != tt.want {
				t.Errorf("GetSelectorFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}
