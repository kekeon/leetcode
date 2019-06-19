package array

func sumEvenAfterQueries(A []int, queries [][]int) []int {

	val := 0
	for _, a := range A {
		if a%2 == 0 {
			val += a
		}
	}

	res := []int{}
	for _, v := range queries {

		if A[v[1]]%2 == 0 {
			val -= A[v[1]]
		}

		A[v[1]] += v[0]
		if A[v[1]]%2 == 0 {
			val += A[v[1]]
		}

		res = append(res, val)
	}

	return res

}
