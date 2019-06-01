package array

func findLengthOfLCIS(nums []int) int {

	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return 1
	}

	currentL := 1
	lastL := 1

	for i := 1; i < l; i++ {
		if nums[i] > nums [i-1] {
			currentL ++
		} else if currentL > lastL {
			lastL = currentL
		} else {
			currentL = 1
		}
	}

	return lastL
}
