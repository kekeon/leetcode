package array

func pivotIndex(nums []int) int {

	l := len(nums)

	if l < 3 {
		return -1
	}
	l_sum := 0
	sum := 0

	for j := 0; j < l; j++ {
		sum+=nums[j]
	}


	for i := 0; i < l; i++ {

		if l_sum * 2 == sum - nums[i] {
			return i
		}

		l_sum += nums[i]
	}

	return -1
}
