package dp

func numberOfArithmeticSlices(A []int) int {

	l := len(A)

	if l < 3 {
		return 0
	}

	before := 0
	after := 0
	sum := 0
	for a := 2; a < l; a++ {
		if A[a]-A[a-1] == A[a-1]-A[a-2] {
			after = before + 1
			before = after
			sum += before
		} else {
			before = 0
			after = 0
		}
	}

	return sum
}
