package array

func findMaxConsecutiveOnes(nums []int) int {
	i := 0
	if len(nums) == 0 {
		return i
	}

	max := 0

	for _, v := range nums {
		if v != 1 {
			if i > max {
				max = i
			}
			i = 0
		}else {
			i ++
		}
	}

	if i > max {
		max = i
	}

	return max

}
