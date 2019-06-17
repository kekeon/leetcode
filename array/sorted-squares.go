package array

import "sort"

func sortedSquares(A []int) []int {
	for a:=0; a< len(A); a++ {
		A[a] = A[a] * A[a]
	}
	sort.Ints(A)
	return A
}
