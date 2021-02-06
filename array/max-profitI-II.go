package array

func maxProfitII(prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := 0
	sum := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > max {
			max = v - min
			sum += max
			max = 0
			min = v
		}
	}

	return sum
}

/**
定义二维数组 dp[n][2]dp[n][2]：
dp[i][0] 表示第 ii 天不持有可获得的最大利润；
dp[i][1] 表示第 ii 天持有可获得的最大利润（注意是第 ii 天持有，而不是第 ii 天买入）。

定义状态转移方程：
不持有：dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])dp[i][0]=max(dp[i−1][0],dp[i−1][1]+prices[i])
	对于今天不持有，可以从两个状态转移过来：1. 昨天也不持有；2. 昨天持有，今天卖出。两者取较大值。
持有：dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])dp[i][1]=max(dp[i−1][1],dp[i−1][0]−prices[i])
	对于今天持有，可以从两个状态转移过来：1. 昨天也持有；2. 昨天不持有，今天买入。两者取较大值
 */
func maxProfitDPII(prices []int) int {
	lp := len(prices)
	if len(prices) == 0 {
		return 0
	}

	arr := make([][2]int, lp)

	arr[0][0] = 0
	arr[0][1] = -prices[0]
	for a := 1; a < lp; a++ {
		arr[a][0] = Max(arr[a - 1][0], arr[a - 1][1] + prices[a])
		arr[a][1] = Max(arr[a - 1][1], arr[a - 1][0] - prices[a])
	}

	return arr[lp-1][0]
}

// diff 做差法
func maxProfitDiffII(prices []int) int {
	lp := len(prices)
	if len(prices) == 0 {
		return 0
	}

	sum := 0
	for a := 1; a < lp; a++ {
		v := prices[a] - prices[a-1]
		if v > 0 {
			sum += v
		}
	}

	return sum
}
