package array

func findPairs(nums []int, k int) int {

	i := 0


	m := map[int]int{}

	arr := []int{}

	for _, val := range nums {
		if m[val] < 2 {
			m[val] ++
			arr = append(arr, val)
		}
	}

	mv := map[int]int{}
	mv1 := map[int]int{}

	for j := 0; j < len(arr) - 1; j++ {
		v := arr[j]
		for _, v1 := range arr[j+1:] {

			if v == 0 && v1 == v && k == 0 {
				i++
			}else if intAbs(v-v1) == k && mv[v] != v1 && mv1[v1] != v  {
				i++
				mv[v] = v1
				mv1[v1] = v
			}
		}
	}

	return i
}
