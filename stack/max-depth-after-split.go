package stack

func maxDepthAfterSplit(seq string) []int {
	dep := 0
	arr := []int{}

	for _, v := range seq {
		if v == '(' {
			dep ++
			arr = append(arr, dep)
		} else {

			arr = arr[0: len(arr) - 1]
			arr = append(arr, dep)
		}
	}

	return arr
}
