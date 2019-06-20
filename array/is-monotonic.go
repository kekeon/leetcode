package array

func isMonotonic(A []int) bool {
	l := len(A)

	if l <= 2 {
		return true
	}

	var tag int

	tag = A[1] - A[0]

	for i := 1; i < l - 1 ; i++ {
		if tag < 0 && A[i] < A[i+1]{
			return false
		}else if tag > 0 && A[i] > A[i+1] {
			return false
		}else if tag == 0 && A[i] != A[i+1]{
			tag = A[i+1] - A[i]
		}
	}

	return true
}