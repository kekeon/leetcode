package hash_table
func findLHS(nums []int) int {

	m := map[int]int{}

	for _, v := range nums {
		m[v]++
	}

	sum, res := 0, 0
	for k, kv := range m {
		if m[k+1] != 0 {
			sum = kv + m[k+1]
			res = max(res, sum)
		}
	}

	return res
}

func max(a,b int) int {

	if a >= b {
		return a
	}

	return b
}
