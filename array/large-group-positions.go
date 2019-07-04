package array

func largeGroupPositions(S string) [][]int {

	res := [][]int{}

	cache, indexs, num := string(S[0]), make([]int, 2), 0
	indexs[0] = 0
	for i, v := range S {

		if cache == string(v) {
			num ++
		}

		cache = v

	}

	return res
}
