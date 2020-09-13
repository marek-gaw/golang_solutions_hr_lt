/*
 *	Site: 	Hacker Rank
 *	Link: 	https://www.hackerrank.com/challenges/repeated-string
 *  Level:	Easy
 */

package repeatedstring

import "strings"

func repeatedString(s string, n int64) int64 {

	var result int = 0
	data := []byte(s)

	repeats := n / int64(len(data))
	modulo := n - (int64(len(data)) * int64(repeats))

	result = int(repeats) * strings.Count(s, "a")

	for _, v := range data[:int(modulo)] {
		if v == 'a' {
			result++
		}
	}

	return int64(result)
}
