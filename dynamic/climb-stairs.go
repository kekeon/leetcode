package dynamic

// 计数型动态规划，转移方程 f[n] = f[n-1] + f[n-2]
func climbStairs(n int) int {

	if n == 0 || n == 1 || n == 2{
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
