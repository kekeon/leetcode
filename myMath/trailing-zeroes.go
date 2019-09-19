package myMath

func trailingZeroes(n int) int {
	count := 0
	for n >=5 {
		n = n / 5
		count += n
	}

	return count
}
