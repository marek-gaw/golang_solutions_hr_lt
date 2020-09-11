package main

import "testing"

func TestCountingValleys(t *testing.T) {

	type Data struct {
		result int32
		input  string
	}

	tables := []Data{
		{1, "UDDDUDUU"},
		{2, "DDUUDDUDUUUD"},
		{0, "UDUUUDUDDD"},
	}

	for _, table := range tables {
		res := countingValleys(int32(len(table.input)), table.input)
		if res != table.result {
			t.Errorf("No. of pairs: expected: %d, got: %d", table.result, res)
		}
	}
}
