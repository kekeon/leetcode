package array

import "fmt"

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
		fmt.Println(path[0:])
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := 0; i < l; i++ {
		if useMap[i] {
			continue
		}
		path = append(path, nums[i])
		useMap[i] = true
		dfsArr(nums, l, depth+1, path, useMap, res)
		// l1 := len(path)
		path = path[0 : len(path)-1]
		useMap[i] = false

	}

}
