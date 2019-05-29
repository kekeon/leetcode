package array

import "sort"

func heightChecker(heights []int) int {

	l := len(heights)

	if l == 1 {
		return 0
	}

	nums := append([]int{}, heights...)
 	sort.Ints(nums)
	count := 0
	for i,v := range heights {
		if v != nums[i] {
			count ++
		}
	}
	return count
}