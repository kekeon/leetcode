package array

func minCostClimbingStairs(cost []int) int {

	first , second , sum := 0, 0, 0

	for i := 2; i <= len(cost); i++ {
		sum = min(first + cost[i - 2], second + cost[i-1])
		first, second = second, sum
	}

	return sum

}

func min(a , b int) int{

	if a >= b {
		return b
	}

	return a
}
