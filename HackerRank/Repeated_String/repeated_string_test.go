package repeatedstring

import "testing"

func TestRepeatedStrings(t *testing.T) {

	type Data struct {
		result int64
		input1 string
		input2 int64
	}

	tables := []Data{
		{7, "aba", 10},
		{1000000000000, "a", 1000000000000},
	}

	for _, table := range tables {
		res := repeatedString(table.input1, table.input2)
		if res != table.result {
			t.Errorf("Expected: %d, got: %d", table.result, res)
		}
	}
}
