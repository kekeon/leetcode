package array

func majorityElement(nums []int) int {

	m := map[int]int{}

	l := float64(len(nums)) / 2
	for _, v := range nums {
		m[v]++
		if float64(m[v]) > l {
			return v
		}
	}

	return 0

}
