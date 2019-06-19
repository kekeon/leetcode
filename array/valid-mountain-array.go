package array

func validMountainArray(A []int) bool {

	l := len(A)

	if l < 3 {
		return false
	}

	i := 0

	// 上山
	for i < l -1 && A[i] < A[i+1] {
		i++
	}

	if i==0 || i == l -1 {
		return false
	}

	// 下山
	for i < l - 1 && A[i] > A[i+1] {
		i++
	}

	if  i != l -1 {
		return false
	}

	return true

}