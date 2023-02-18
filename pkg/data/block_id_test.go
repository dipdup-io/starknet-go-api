package data

import (
	"testing"
)

func Test_BlockFilterValidation(t *testing.T) {

	goodTests := []BlockID{
		{Hash: "test"},
		{Number: 1},
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
		{Hash: "test", Number: 1},
		{Hash: "test", String: Latest},
		{Hash: "test", String: Pending},
		{Number: 1, String: Latest},
		{Number: 1, String: Pending},
		{Hash: "test", Number: 1, String: Latest},
		{Hash: "test", Number: 1, String: Pending},
		{String: "invalid"},
	}

	for i := range badTests {
		if err := badTests[i].Validate(); err == nil {
			t.Errorf("[%d] invalid struct in bad: %v", i, badTests[i])
		}
	}
}
