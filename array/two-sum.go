package array

func twoSum(numbers []int, target int) []int {
	mIdx := 0
	l := len(numbers)
	lIdex := l - 1
	v := 0
	for {
		v = numbers[mIdx] + numbers[lIdex]
		if v == target {
			break
		} else if v < target {
			mIdx ++
		} else {
			lIdex --
		}
	}
	return []int{mIdx + 1, lIdex + 1}
}
