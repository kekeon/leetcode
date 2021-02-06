package dp

import "leetcode/array"

func maxProfit(prices []int) int {
	lp := len(prices)
	if len(prices) == 0 {
		return 0
	}

	arr := make([][2]int, lp)

	arr[0][0] = 0
	arr[0][1] = -prices[0]
	for a := 1; a < lp; a++ {
		arr[a][0] = array.Max(arr[a - 1][0], arr[a - 1][1] + prices[a])
		arr[a][1] = array.Max(arr[a - 1][1], arr[a - 1][0] - prices[a])
	}

	return arr[lp-1][0] + arr[lp-2][0]
}
