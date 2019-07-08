package array

func largeGroupPositions(S string) [][]int {

	res := [][]int{}

	l := len(S)

	if l < 3 {
		return res
	}

	a, b, c := 1, 0, 0

	for ; a < l; a++ {
		if S[b] != S[a] {
			if c != 0 {
				list := []int{}
				if c >= 2 {
					list = append(list, b)
					list = append(list, a-1)
					res = append(res, list)
				}

				b = a - 1
				c = 0
			}
			b ++
		} else {
			c ++
		}
	}

	if c >= 2 {
		list := []int{}
		list = append(list, b)
		list = append(list, a-1)
		res = append(res, list)
	}

	return res
}
