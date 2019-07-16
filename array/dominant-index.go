package array

func dominantIndex(nums []int) int {

	max, subMax := 0, 0

	if len(nums) == 1 {
		return 0
	}

	j := 0
	for i, v := range nums {
		if v > max {
			subMax = max
			max = v
			j = i
		} else if v > subMax {
			subMax = v
		}
	}

	if subMax*2 <= max {
		return j
	}

	return -1

}
