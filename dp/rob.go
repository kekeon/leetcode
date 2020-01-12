package dp

/**
198. 打家劫舍
1. 子问题
2. 确定状态
3. 动态转移方程

dp[a][0] = max(dp[a-1][1], dp[a-1][0])
dp[a][1] = dp[a-1][0] + nums[a]
 */
func rob(nums []int) int {

	l := len(nums)
	if len(nums) < 1 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}


	dp := make([][]int,l)
	dp[0] = make([]int, 2)
	// dp[n][0] 不偷, dp[n][1] 偷
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for a := 1; a < l; a++ {
		dp[a] = make([]int, 2)
		dp[a][0] = max(dp[a-1][1], dp[a-1][0])
		dp[a][1] = dp[a-1][0] + nums[a]
	}

	return max(dp[l-1][0], dp[l-1][1])
}


// 转移方程 res := max(dp[i-1] + 0, dp[i-2] + dp[i])
func rob1(nums []int) int {

	l := len(nums)
	if len(nums) < 1 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	if l == 2 {
		return max(nums[0], nums[1])
	}


	dp := make([]int,l)
	dp[0] = nums[0]
	dp[1] = nums[1]

	for a := 2; a < l; a++ {
		dp[a] = max(dp[a-1], nums[a-2] + nums[a])
	}

	return dp[l-1]
}
