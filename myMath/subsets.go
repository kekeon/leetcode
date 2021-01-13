package myMath

import "sort"

func subsets(nums []int) [][]int {
	//result := [][]int{}

	// subset(&result, nums, []int{}, len(nums) - 1)
	return forsubset(nums)
	// return result
}

func subsetRecursion(result *[][]int, nums []int, arr []int, index int) {

}

func subset(result *[][]int, nums []int, arr []int, index int) {

	if index == -1 {
		*result = append(*result, arr)
		return
	}

	subset(result, nums, arr, index - 1)
	arr = append(arr, nums[index])

	subset(result, nums, arr, index - 1)

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
