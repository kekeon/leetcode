package array

func maxSubArray(nums []int) int {
	l := len(nums)

	if l == 1 {
		return nums[0]
	}

	sum  := 0
	maxSum := nums[0]

	for i:=0; i<l; i++ {
		if sum > 0 {
			sum += nums[i]
		}else {
			sum = nums[i]
		}
		if maxSum < sum {
			maxSum = sum
		}
	}

	return maxSum

}