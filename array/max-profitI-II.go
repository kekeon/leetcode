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