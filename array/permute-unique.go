package array

import (
	"strconv"
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {

	l := len(nums)
	res := [][]int{}

	if l == 0 {
		return res
	}

	path := []int{}
	useMap := map[int]bool{}
	inRes := map[string]bool{}


	defUniqueArr(nums, l, path, 0, useMap, "", inRes, &res)

	path2 := []int{}
	useMap2 := map[int]bool{}
	res2 := [][]int{}
	// 剪枝前要排序
	sort.Ints(nums)
	defUnique(nums, l, path2, 0, useMap2, &res2)

	fmt.Println(len(res), len(res2))

	return res

}

func defUniqueArr(nums []int, l int, path []int, depth int, useMap map[int]bool, pathStr string, inRes map[string]bool, res *[][]int) {

	if l == depth {
		if !inRes[pathStr] {
			inRes[pathStr] = true
			*res = append(*res, append([]int{}, path...))
		}
		return
	}

	for i := 0; i < l; i++ {

		if useMap[i] {
			continue
		}
		useMap[i] = true
		path = append(path, nums[i])
		numIs := strconv.Itoa(nums[i])
		pathStr += numIs
		defUniqueArr(nums, l, path, depth+1, useMap, pathStr, inRes, res)
		useMap[i] = false
		path = path[0 : len(path)-1]
		pathStr = pathStr[0 : len(pathStr)-len(numIs)]

	}

}

func defUnique(nums []int, l int, path []int, depth int, useMap map[int]bool, res *[][]int) {

	if l == depth {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := 0; i < l; i++ {

		if useMap[i] {
			continue
		}

		// 剪支条件
		if i > 0 && nums[i] == nums[i-1] && useMap[i-1] == false {
			continue
		}

		useMap[i] = true
		path = append(path, nums[i])
		defUnique(nums, l, path, depth+1, useMap, res)
		useMap[i] = false
		path = path[0 : len(path)-1]

	}

}
