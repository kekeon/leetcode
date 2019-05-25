package array

import (
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)

	if nums[0] > 0 || nums[l-1] <= 0 {
		return nums[l-3] * nums[l-2] * nums[l-1]
	}else  if nums[0] * nums[1] > nums[l-3] * nums[l-2] {
		return nums[0] * nums[1] * nums[l-1]
	}else {
		return nums[l-3] * nums[l-2] * nums[l-1]
	}
}