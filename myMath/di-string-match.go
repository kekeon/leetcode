package myMath

func diStringMatch(S string) []int {

	start := 0
	end := len(S)

	arr := []int{}

	for _,v := range S {
		if string(v) == "I" {
			arr = append(arr, start)
			start ++
		}else {
			arr = append(arr, end)
			end --
		}
	}

	arr = append(arr, end)

	return arr
}
