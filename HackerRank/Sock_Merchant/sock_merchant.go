/*
 *	Site: 	Hacker Rank
 *	Link: 	https://www.hackerrank.com/challenges/sock-merchant/problem
 *  Level:	Easy
 */

package sockmerchant

import (
	"sort"
)

func SockMerchant(n int32, ar []int32) int32 {
	sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })

	var i, pairs int32

	for i < n-1 {
		if ar[i] == ar[i+1] {
			pairs++
			i++
		}
		i++
	}

	return pairs
}
