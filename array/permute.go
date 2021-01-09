package array

func permute(nums []int) [][]int {
	res := [][]int{}
	l := len(nums)
	if l == 0 {
		return res
	}
	path := []int{}
	useMap := map[int]bool{}

	dfsArr(nums, l, 0, path, useMap, &res)

	return res
}

func dfsArr(nums []int, l, depth int, path []int, useMap map[int]bool, res *[][]int) {
	if depth == l {
		*res = append(*res, path)
		return
	}

	for i := 0; i < l; i++ {
		if useMap[i] {
			continue
		}
		path = append(path, nums[i])
		useMap[i] = true
		dfsArr(nums, l, depth+1, path, useMap, res)
		l1 := len(path)
		path = path[0 : l1-1]
		useMap[i] = false

	}

}
