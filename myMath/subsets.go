package myMath

func subsets(nums []int) [][]int {
	result := [][]int{}

	subset(&result, nums, []int{}, len(nums) - 1)

	return result
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
	result = append(result, []int{})

	if len(nums) == 0 {
		return result
	}

	for _, v := range nums {
		arr := [][]int{}
		for _, v2 := range result {
			n := []int{}
			v2 = append(v2, v)
			n = append(n, v2...)
			arr = append(arr, n)
		}

		result = append(result, arr...)
	}

	return result
}
