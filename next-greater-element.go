package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	res := []int{}

	i := 0

	b := -1
	for _, v := range nums1 {

		i = fincIndex(nums2, v)
		b = -1
		for a := i; a < len(nums2); a++ {
			if nums2[a] > v {
				b = nums2[a]
				break
			}
		}

		res = append(res, b)
	}

	return res

}

func fincIndex(s []int, num int) int {

	for i, v := range s {

		if v == num {
			return i
		}
	}

	return -1
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	v := nextGreaterElement(nums1, nums2)
	fmt.Println(v)
}
