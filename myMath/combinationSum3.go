package myMath

import "fmt"

func combinationSum3(k int, n int) [][]int {

	result := [][]int{}
	combinationRecursionSum3(n, 1, k, map[int]bool{}, []int{}, &result)
	return result
}

func combinationRecursionSum3(target, start, depth int, useMap map[int]bool, path []int, result *[][]int) {

	// 终止条件

	if target < 0 || (target > 0 && depth <= len(path)) {
		return
	}

	if target == 0 && depth == len(path) {
		*result = append(*result, append([]int{}, path...))
		return
	}

	max := target
	if max >= 10 {
		max = 10
	}

	for a := start; a <= max; a++ {

		// 剪枝
		if useMap[a] {
			continue
		}
		useMap[a] = true
		path = append(path, a)
		fmt.Println("start===========>", path, a)
		combinationRecursionSum3(target-a, a+1, depth, useMap, path, result)
		useMap[a] = false
		path = path[0 : len(path)-1]
		fmt.Println("end===========>", path, target)
	}

}
