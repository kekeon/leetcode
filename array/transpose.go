package array

func transpose(A [][]int) [][]int {


	if len(A) == 1 && len(A[0]) == 1 {
		return A
	}

	res := [][]int{}
	for j := 0; j < len(A[0]); j++ {
		a := []int{}
		for i := 0; i < len(A); i ++ {
			a = append(a, A[i][j])
		}
		res = append(res, a)
	}
	return res

}
