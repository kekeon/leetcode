package array

import "fmt"

func sumEvenAfterQueries(A []int, queries [][]int) []int {


	val := 0
	for _, a := range A {
		if EvenNumber(a) {
			val += a
		}
	}

	res := []int{}


	for _,v := range queries {
		A[v[1]] += v[0]

		fmt.Println(A[v[1]])
		if EvenNumber(A[v[1]]) {

			if A[v[1]] < 0 {
				val += A[v[1]] - (A[v[1]] - v[0])
			}else {
				val += A[v[1]]
			}

			res = append(res, val)
		}else {
			val -= (A[v[1]] - v[0])
			res = append(res, val)
		}
	}

	return res

}

func EvenNumber(n int) bool{
	if (n >= 0 && n % 2 == 0) || (n < 0 && -n % 2 == 0)  {
		return true
	}

	return false
}