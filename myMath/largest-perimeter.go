package myMath

import "sort"

// 两边之和大于第三边， 两边之差小于第三边
// c < b + a, c > b - a
func largestPerimeter(A []int) int {

	sort.Ints(A)
	for a := len(A) - 1; a >= 2; a-- {
		if A[a] < (A[a-1] + A[a-2]) {
			return A[a] + A[a-1] + A[a-2]
		}
	}

	return 0
}
