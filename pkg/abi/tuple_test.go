package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractTupleTypes(t *testing.T) {
	tests := []struct {
		name    string
		typ     string
		want    []tupleItem
		wantErr bool
	}{
		{
			name: "test 1",
			typ:  "(felt, b: felt, c: felt)",
			want: []tupleItem{
				{
					Type: Type{
						Name: "tuple_item_0",
						Type: "felt",
					},
				}, {
					Type: Type{
						Name: "b",
						Type: "felt",
					},
				}, {
					Type: Type{
						Name: "c",
						Type: "felt",
					},
				},
			},
		}, {
			name: "test 2",
			typ:  "(d: felt, (felt, b: felt, (f: felt, g:felt)))",
			want: []tupleItem{
				{
					Type: Type{
						Type: "felt",
						Name: "d",
					},
				},
				{
					Type: Type{
						Type: "(felt,b:felt,(f:felt,g:felt))",
						Name: "tuple_item_1",
					},
					Childs: []tupleItem{
						{
							Type: Type{
								Name: "tuple_item_0",
								Type: "felt",
							},
						}, {
							Type: Type{
								Name: "b",
								Type: "felt",
							},
						}, {
							Type: Type{
								Name: "tuple_item_2",
								Type: "(f:felt,g:felt)",
							},
							Childs: []tupleItem{
								{
									Type: Type{
										Name: "f",
										Type: "felt",
									},
								}, {
									Type: Type{
										Name: "g",
										Type: "felt",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractTupleTypes(tt.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractTupleTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "items")
		})
	}
}
