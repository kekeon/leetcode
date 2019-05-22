package array

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0

	l := len(nums)
	for i := 0; i < l; i+=2 {
		sum += nums[i]
	}

	return sum
}
