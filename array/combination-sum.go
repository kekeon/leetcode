package array

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {

	l := len(candidates)
	res := [][]int{}

	path := []int{}

	combinationRecursion(candidates, l, target, 0, path, &res)

	res2 := [][]int{}

	path2 := []int{}

	sort.Ints(candidates)
	combinationRecursion2(candidates, l, target, 0, path2, &res2)

	fmt.Println("======", res, res2)

	return res

}

func combinationRecursion(nums []int, l, target, start int, path []int, res *[][]int) {

	// 终止条件
	if target < 0 {
		return
	}

	// 结果满足
	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := start; i < l; i++ {

		path = append(path, nums[i])
		// 递归阶段
		fmt.Println("=====>递归前", path)
		combinationRecursion(nums, l, target-nums[i], i, path, res)
		// 回溯
		path = path[0 : len(path)-1]
		fmt.Println("=====>递归后", path)

	}

}

func combinationRecursion2(nums []int, l, target, start int, path []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := start; i < l; i++ {
		if target-nums[i] < 0 {
			break
		}

		path = append(path, nums[i])
		fmt.Println("=====>递归前", path)
		combinationRecursionUnique(nums, l, target-nums[i], i, path, res)
		path = path[0 : len(path)-1]
		fmt.Println("=====>递归后", path)

	}

}
