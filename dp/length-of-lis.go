package dp

/**
300. 最长上升子序列
[10,9,2,5,3,7,101,18] -> 4
 */
func lengthOfLIS(nums []int) int {
	l := len(nums)

	if l <= 1 {
		return l
	}

	arr := make([]int, l)
	arr[0] = 1
	maxCount:= 1

	for i := 1; i < l; i++{

		maxVal := 0
		for j := 0; j < l; j++ {

			if nums[i] > nums[j] {
				maxVal = max(maxVal, arr[j])
			}
		}
		arr[i] = maxVal + 1
		maxCount = max(arr[i], maxCount)

	}


	return maxCount

}

func max(a, b int) int{
	if a > b {
		return a
	}

	return b
}
