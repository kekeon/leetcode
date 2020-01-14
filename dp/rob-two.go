package dp

/**
213
 */
func robTwo(nums []int) int {

	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	if l == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([][]int, l-1)
	dp[1] = make([]int, 2)
	dp[0] = make([]int, 2)

	dp1 := make([][]int, l-1)
	dp1[1] = make([]int, 2)
	dp1[0] = make([]int, 2)

	// 1 偷， 0 不偷
	dp[0][1] = nums[0]
	dp[0][0] = 0

	dp1[0][1] = nums[1]
	dp1[0][0] = 0

	for a := 1; a < l-1; a++ {
		dp[a] = make([]int, 2)
		dp[a][1] = dp[a-1][0] + nums[a]
		dp[a][0] = max(dp[a-1][1], dp[a-1][0])
	}

	for b := 2; b < l; b++ {
		dp1[b-1] = make([]int, 2)
		dp1[b-1][1] = dp1[b-2][0] + nums[b]
		dp1[b-1][0] = max(dp1[b-2][1], dp1[b-2][0])
	}

	return max(max(dp[l-2][1], dp[l-2][0]), max(dp1[l-2][1], dp1[l-2][0]))
}

func robTwo2(nums []int) int {

	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	if l == 2 {
		return max(nums[0], nums[1])
	}

	return max(robTwoMax2(0, l-2, nums), robTwoMax2(1, l-1, nums))
}

func robTwoMax(start int, end int, nums []int) int {

	l := end - start + 1

	arr := make([]int, l)
	arr[0] = nums[start]
	arr[1] = max(nums[start], nums[start+1])

	for i := 2; i < l; i++ {
		arr[i] = max(arr[i-1], arr[i-2]+nums[start+i])
	}

	return arr[l-1]
}

func robTwoMax2(start int, end int, nums []int) int {

	l := end - start + 1

	after := nums[start]
	before := max(nums[start], nums[start+1])
	res := max(after, before)
	for i := 2; i < l; i++ {
		res = max(before, after+nums[start+i])
		after = before
		before = res
	}

	return res
}
