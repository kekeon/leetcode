package array

func addToArrayForm(A []int, K int) []int {

	l := len(A)

	arr := []int{}

	current := K
	for i := l - 1; i >= 0 || current > 0; i-- {

		if i >= 0 {
			current += A[i]
		}

		arr = append(arr, current%10)
		current /= 10
	}

	l = len(arr)
	current = l / 2

	for j:=0;j<current;j++ {
		arr[l-j-1],arr[j] = arr[j], arr[l-j-1]
	}

	return arr

}
