package array

func addToArrayForm(A []int, K int) []int {

	l := len(A)

	arr := []int{}

	current := K
	for i := l - 1; i >= 0 || current > 0; i-- {

		if i >= 0 {
			current += A[i]
		}

		c := append([]int{}, current%10)
		arr = append(c, arr...)
		current /= 10
	}

	return arr

}
