package dp

import (
	"math"
	"fmt"
)

func maxProduct(nums []int) int {

	l := len(nums)

	if l <= 1 {
		return nums[0]
	}

	maxNum := math.MinInt8
	tag := math.MinInt8

	for i := 0; i <= l-1; i++ {

		num := nums[i]

		for j := i + 1; j < len(nums); j++ {
			tag = num * nums[j]
			if tag < nums[j] {
				break
			}
			num = tag
		}

		nums[i] = num
		fmt.Println(maxNum, num)
		maxNum = max(maxNum, num)
	}

	return maxNum
}
