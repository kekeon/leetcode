package array

import (
	"math"
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := []int{}
	l1 := len(nums1)
	l2 := len(nums2)
	i1 := 0
	i2 := 0

	for true {

		v1 := math.MinInt64
		v2 := math.MinInt64

		if l1 > i1 {
			v1 = nums1[i1]
		}

		if l2 > i2 {
			v2 = nums2[i2]
		}

		switch {
		// v2 小，且v2 有效范围
		case v1 > v2 && l2 > i2:
			arr = append(arr, v2)
			i2++
			// v1 小，且v1 有效范围
		case v1 < v2 && l1 > i1:
			arr = append(arr, v1)
			i1++
			// v1 小，且v1 不在效范围
		case v1 < v2 && l1 <= i1:
			arr = append(arr, nums2[i2:]...)
			goto END
		case v1 > v2 && l2 <= i2:
			arr = append(arr, nums1[i1:]...)
			goto END
			// 两个相等，有效范围
		case v1 == v2 && l2 > i2 && l1 > i1:
			arr = append(arr, v1)
			arr = append(arr, v1)
			i1++
			i2++
			// 两个相等，无效范围
		case v1 == v2 && l2 <= i2 && l1 <= i1:
			goto END
		}
	}
END:
	fmt.Println(arr)
	l := l1 + l2
	if l%2 == 0 {
		c := l / 2
		return float64(arr[c] + arr[c-1]) / 2
	} else {
		c := l / 2
		return float64(arr[c])
	}
	return 0

}
