package array

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := 0

	for _, v := range prices {
		if v < min {
			min = v
		}else if v - min > max {
			max = v - min
		}
	}

	return max
}
