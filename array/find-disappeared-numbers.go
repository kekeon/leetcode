package array

func findDisappearedNumbers(nums []int) []int {


	l := len(nums)

	if l == 0 {
		return nums
	}

	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	arr := []int{}
	for i := 1; i < l; i++ {
		if m[i] == 0 {
			arr = append(arr, i)
		}
	}

	return arr

}
