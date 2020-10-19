package RemoveDuplicatesFromSortedArray

import "sort"

func removeDuplicates(nums []int) int {
	retval := len(nums)

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			retval--
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return retval
}
