// inspired by: https://blog.alexellis.io/golang-writing-unit-tests/
package sockmerchant

import "testing"

func TestSockMerchant(t *testing.T) {

	type Data struct {
		result int32
		input  []int32
	}

	tables := []Data{
		{2, []int32{1, 1, 2, 2}},
		{3, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}},
	}

	for _, table := range tables {
		pairs := SockMerchant(int32(len(table.input)), table.input)
		if pairs != table.result {
			t.Errorf("No. of pairs: expected: %d, got: %d", table.result, pairs)
		}
	}
}
