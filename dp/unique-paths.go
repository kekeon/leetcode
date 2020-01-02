package dp

/**
62. 不同路径

计数型动态规划，转移方程 f[n][m] = f[n-1][m] + f[n][m-1]
 */
func uniquePaths(m int, n int) int {

	if m == 0 && n == 0 {
		return 0
	}

	if (m == 1 && n == 0) || (m == 0 && n == 1) {
		return 1
	}

	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	arr[0][0] = 0
	for j := 0; j < n; j++ {
		for t := 0; t < m; t++ {
			if j == 0 || t == 0 {
				arr[j][t] = 1
			} else if j-1 >= 0 && t-1 >= 0 {
				arr[j][t] = arr[j-1][t] + arr[j][t-1]
			}
		}
	}

	return arr[n-1][m-1]
}
