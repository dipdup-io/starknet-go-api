package data

import (
	"testing"
)

func Test_BlockFilterValidation(t *testing.T) {
	number := new(uint64)
	*number = 1
	goodTests := []BlockID{
		{Hash: "test"},
		{Number: number},
		{String: Latest},
		{String: Pending},
	}

	for i := range goodTests {
		if err := goodTests[i].Validate(); err != nil {
			t.Errorf("[%d] invalid struct in good: %v %s", i, goodTests[i], err)
		}
	}

	badTests := []BlockID{
		{},
		{Hash: "test", Number: number},
		{Hash: "test", String: Latest},
		{Hash: "test", String: Pending},
		{Number: number, String: Latest},
		{Number: number, String: Pending},
		{Hash: "test", Number: number, String: Latest},
		{Hash: "test", Number: number, String: Pending},
		{String: "invalid"},
	}

	for i := range badTests {
		if err := badTests[i].Validate(); err == nil {
			t.Errorf("[%d] invalid struct in bad: %v", i, badTests[i])
		}
	}
}
