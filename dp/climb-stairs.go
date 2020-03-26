package dp

/**
70. 爬楼梯
计数型动态规划，转移方程 f[n] = f[n-1] + f[n-2]
 */
func climbStairs(n int) int {

	if n <= 2{
		return n
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2

	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

func climbStairsOptimization(n int) int {

	if n  <= 2{
		return n
	}

	a := 1
	b := 2
	c := 3

	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b =c
	}

	return c
}
