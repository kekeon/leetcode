package myMath

func grayCode(n int) []int {
	arr := []int{0}

	if n <= 0 {
		return arr
	}

	for a := 1; a < 1<<uint(n); a++ {
		arr = append(arr, a ^ a>>1)
	}
	return arr

}
