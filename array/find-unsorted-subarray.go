package array

import "sort"

func findUnsortedSubarray(nums []int) int {
	l := len(nums)

	if l == 1 {
		return 0
	}

	arr := append([]int{}, nums...)
	sort.Ints(nums)
	i := 0
	j := l - 1

	for i<j && nums[i] == arr[i] {
		i++
	}

	for i<=j && nums[j] == arr[j] {
		j--
	}

	return j - i + 1
}
