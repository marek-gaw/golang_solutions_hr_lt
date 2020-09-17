package equalizethearray

import "testing"

func TestEqualizeArray(t *testing.T) {

	type Data struct {
		result int32
		input  []int32
	}

	tables := []Data{
		{result: 2, input: []int32{3, 3, 2, 1, 3}},
		{result: 4, input: []int32{1, 2, 3, 1, 2, 3, 3, 3}},
		{result: 5, input: []int32{37, 32, 97, 35, 76, 62}},
	}

	for _, table := range tables {
		res := equalizeArray(table.input)
		if res != table.result {
			t.Errorf("No. of pairs: expected: %d, got: %d", table.result, res)
		}
	}
}
