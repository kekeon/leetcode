package hashTable

func intersect(nums1 []int, nums2 []int) []int {

	res := []int{}
	m := map[int]int{}

	for _, v1 := range nums1 {
		m[v1]++
	}
	for _, v2 := range nums2 {
		if m[v2] > 0 {
			res = append(res, v2)
			m[v2]--
		}
	}
	return res
}
