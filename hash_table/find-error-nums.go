package hash_table

// 645. Set Mismatch， 等差数列求解 s(N) = Na1 + N(N-1)d/2
// 差 =  数列总和 - 当前总和
func findErrorNums(nums []int) []int {

	l := len(nums)
	sum := l + (l * l - l)/2

	count := 0
	m := map[int]int{}

	lab := 0

	for _, v := range nums {
		if m[v] == 0 {
			m[v] ++
			count += v
		} else {
			lab = v
		}
	}

	return []int{lab, sum - count}

}
