package myMath

// 892
func surfaceArea(grid [][]int) int {

	total := 0

	col := len(grid)
	for i := 0; i < col; i++ {
		row := len(grid[i])

		// 第一列，或者列中有间隔0
		if grid[i][0] > 0 {
			if i == 0 || grid[i-1][0] == 0 {
				total += count(grid[i][0])
			} else {
				total += count(grid[i][0]) - min( grid[i][0],  grid[i-1][0]) * 2
			}
		}

		for j := 1; j < row; j++ {
			if grid[i][j] > 0 {
				total += count(grid[i][j])
				// 周边都有的方块判断
				if i > 0 {

					if grid[i][j-1] > 0 {
						total -= min(grid[i][j-1], grid[i][j]) * 2
					}

					if grid[i-1][j] > 0 {
						total -= min(grid[i-1][j], grid[i][j]) * 2
					}

				} else {
					if  grid[i][j-1] > 0 {
						total -= min(grid[i][j-1], grid[i][j]) * 2
					}
				}

			}
		}
	}

	return total
}

func count(a int) int {
	return (a * 6) - (a-1)*2
}

func min (a, b int) int {
	if a >= b {
		return b
	}

	return a
}
