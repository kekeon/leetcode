package array

import "fmt"

func commonChars(A []string) []string {
	res := []string{}

	resM := map[string]int{}

	for _, v := range A[0] {
		resM[string(v)] ++
	}

	fmt.Println(resM)


	for _, s := range A {
		m := map[string]int{}
		for _, v := range s {

			m[string(v)]++

		}
	}

	return res
}
