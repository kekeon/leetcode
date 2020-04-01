package stack

func maxDepthAfterSplit(seq string) []int {
	arr := []int{}

	for i, v := range seq {
		if v == '(' {
			arr = append(arr, i % 2)
		} else {

			arr = append(arr, 1 - i % 2)
		}
	}

	return arr
}
