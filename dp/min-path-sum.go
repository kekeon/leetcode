package dp

/**
64. 最小路径和
最值动态规划，转移方程 f[i][j] = grid[i][j] + min(grid[i+1][j], grid[i][j+1])

可以使用同一个grid进行逆向推算，使用O(1)的空间
 */

func minPathSum(grid [][]int) int {
	row := len(grid) - 1
	col := len(grid[0]) - 1

	t := 0
	for i := row; i >= 0; i-- {
		for j := col; j >= 0; j-- {

			if i == row && j == col {
				continue
			}

			t = grid[i][j]
			if i == row {
				grid[i][j] = t + grid[i][j+1]
			} else if j == col {
				grid[i][j] = t + grid[i+1][j]
			} else {
				grid[i][j] = t + Min(grid[i+1][j], grid[i][j+1])
			}

		}
	}

	return grid[0][0]
}

func Min(a, b int) int {

	if a < b {
		return a
	}

	return b
}
