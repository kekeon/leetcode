package hash_table

func intersection(nums1 []int, nums2 []int) []int {

	res := []int{}
	m := map[int]int{}

	for _, v1 := range nums1 {
		m[v1] = 1
	}
	for _, v2 := range nums2 {
		if m[v2] == 1 {
			res = append(res, v2)
			m[v2] = 0
		}
	}
	return res
}
