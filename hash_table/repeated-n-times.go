package hash_table

func repeatedNTimes(A []int) int {
	m := map[int]int{}

	for _, v := range A {
		if m[v]++; m[v] >= 2{
			return m[v]
		}

	}

	return 0
}
