package myMath

func binaryGap(N int) int {

	last := -1
	num := 0
	for i := 0;N > 0; i++ {
		if N & 1 > 0 {
			if last >= 0 {
				num = max(num, i-last)
			}
			last = i
		}
		N >>=1
	}

	return num
}

func max(a,b int) int {

	if a >= b {
		return a
	}

	return b
}
