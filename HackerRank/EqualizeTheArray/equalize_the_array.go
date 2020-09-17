/*
 *	Site: 	Hacker Rank
 *	Link: 	https://www.hackerrank.com/challenges/equality-in-a-array
 *  Level:	Easy
 */

package equalizethearray

import "sort"

// Complete the equalizeArray function below.
func equalizeArray(arr []int32) int32 {

	repeats := make(map[int32]int)
	max := 1

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {

			if _, exist := repeats[arr[i]]; exist {

				repeats[arr[i]]++
			} else {
				repeats[arr[i]] = 2
			}
		}
	}

	//find max value in map
	for _, v := range repeats {
		if v > max {
			max = v
		}
	}

	return int32(len(arr) - (max))
}
