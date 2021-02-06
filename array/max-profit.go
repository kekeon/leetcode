package array

import "leetcode/dp"

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := 0

	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > max {
			max = v - min
		}
	}

	return max
}

func maxProfitDP(prices []int) int {

	lp := len(prices)
	if  lp == 0 {
		return 0
	}

	arr := make([]int, lp)
	min := prices[0]
	for a := 1; a < lp; a++ {
		min = dp.Min(min, prices[a])
		arr[a] = Max(arr[a - 1], prices[a] - min)
	}

	return arr[lp-1]
}

func Min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
