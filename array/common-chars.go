package array

func commonChars(A []string) []string {
	res := []string{}

	m:= map[string]int{}
	for _, s := range A {

		for _, v := range s {
			m[string(v)]++
			if m[string(v)] > 3 {
				res = append(res, string(v))
				m[string(v)] = 0
			}
		}
	}

	return res
}