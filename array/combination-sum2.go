package array

import (
	"sort"
	"fmt"
)

func combinationSum2(candidates []int, target int) [][]int {
	l := len(candidates)
	res := [][]int{}
	path := []int{}

	sort.Ints(candidates)
	combinationRecursionUnique(candidates, l, target, 0, path, &res)

	return res
}

func combinationRecursionUnique(nums []int, l int, target int, start int, path []int, res *[][]int) {

	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := start; i < l; i++ {
		if target-nums[i] < 0 {
			break
		}

		if i > start && nums[i] == nums[i- 1] {
			continue
		}

		fmt.Println("=====>递归前", path, nums, i)

		path = append(path, nums[i])
		combinationRecursionUnique(nums, l, target-nums[i], i + 1, path, res)
		path = path[0 : len(path)-1]
		fmt.Println("=====>递归后", path)

	}

}
