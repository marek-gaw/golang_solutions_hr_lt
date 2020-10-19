package RemoveDuplicatesFromSortedArray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type Data struct {
		result int
		input  []int
	}

	tables := []Data{
		{2, []int{1, 1, 2}},
	}

	for _, table := range tables {
		result := removeDuplicates(table.input)
		if result != table.result {
			t.Errorf("RemoveDuplicates: expected len: %d, got: %d", table.result, result)
		}
	}
}
