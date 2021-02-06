package myMath

import "fmt"

func combinationSum3(k int, n int) [][]int {

	result := [][]int{}
	combinationRecursionSum3(n, 1, k,  []int{}, &result)
	return result
}

func combinationRecursionSum3(target, start, depth int, path []int, result *[][]int) {

	// 终止条件

	if target < 0 || (target > 0 && depth <= len(path)) {
		return
	}

	if target == 0 && depth == len(path) {
		*result = append(*result, append([]int{}, path...))
		return
	}

	// 剪枝，不必要每次都从1-9中选取，只要选取比目标数小的数即可
	max := target
	if max > 9 {
		max = 9
	}

	for a := start; a <= max; a++ {
		path = append(path, a)
		fmt.Println("start===========>", path, target)
		combinationRecursionSum3(target-a, a+1, depth, path, result)
		path = path[0 : len(path)-1]
		fmt.Println("end===========>", path, target)
	}

}
