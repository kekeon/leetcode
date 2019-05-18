package array

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}

	for i, v := range nums {

		if m[v] == 0 {
			m[v] = i + 1
		}else if m[v] > 0 && intAbs(i - m[v] + 1 ) <= k {
			return true
		}else {
			m[v] = i + 1
		}

	}

	return false
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}

	return i

}
