package myMath

import (
	"sort"
)

func subsets(nums []int) [][]int {
	return backtrackResultUse(nums)
	// return result
}

func subsetRecursion(result *[][]int, nums []int, arr []int, index int) {

}

func subset(result *[][]int, nums []int, arr []int, index int) {

	if index == -1 {
		*result = append(*result, arr)
		return
	}

	subset(result, nums, arr, index-1)
	arr = append(arr, nums[index])

	subset(result, nums, arr, index-1)

}

func forsubset(nums []int) [][]int {

	result := [][]int{}
	v1 := []int{}
	result = append(result, v1)

	if len(nums) == 0 {
		return result
	}

	for _, v := range nums {
		resultCopy := [][]int{}
		resultCopy = append(resultCopy, result...)
		for j, v2 := range resultCopy {

			v2 = append(v2, v)
			resultCopy[j] = append([]int{}, v2...)
			sort.Ints(resultCopy[j])
		}

		result = append(result, resultCopy...)
	}

	return result
}

func backtrackResult(nums []int) [][]int {
	result := [][]int{}
	for a := 0; a <= len(nums); a++ {
		backtrack(a, 0, []int{}, nums, &result)
	}

	return result
}

func backtrack(size int, start int, path []int, nums []int, result *[][]int) {
	if size == len(path) {
		p := append([]int{}, path...)
		*result = append(*result, p)
		return
	}

	for a := start; a < len(nums); a++ {
		path = append(path, nums[a])
		backtrack(size, a+1, path, nums, result)
		path = path[0 : len(path)-1]
	}
}

func backtrackResultUse(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for a := 0; a <= len(nums); a++ {
		backtrackUse(a, 0, []int{}, nums, &result)
	}

	return result
}

func backtrackUse(size int, start int, path []int, nums []int, result *[][]int) {
	if size == len(path) {
		p := append([]int{}, path...)
		*result = append(*result, p)
		return
	}

	for a := start; a < len(nums); a++ {
		if a > start && nums[a] == nums[a - 1] {
			continue
		}

		path = append(path, nums[a])
		backtrackUse(size, a+1, path, nums, result)
		path = path[0 : len(path)-1]

	}
}
