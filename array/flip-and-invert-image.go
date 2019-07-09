package array

func flipAndInvertImage(A [][]int) [][]int {

	l := len(A[0])

	if l == 0  {
		return A
	}

	if l == 1 {
		A[0][0] = A[0][0]^1
		return A
	}

	var c int

	if l % 2 == 0{
		c = len(A[0]) / 2
	} else {
		c = len(A[0])  / 2 + 1
	}


	for i:=0; i<len(A); i++ {
		for j:=0; j< c; j++ {
			A[i][j], A[i][l-j -1 ] = A[i][l-j - 1]^1, A[i][j]^1
		}
	}

	return A
}
