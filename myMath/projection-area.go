package myMath

func projectionArea(grid [][]int) int {
	d := 0

	for i, arr := range grid {

		maxRow := 0
		maxCol := 0

		for j, v := range arr {
			if v > 0 {
				d++
			}

			maxRow = max(maxRow, arr[j])
			maxCol = max(maxCol, grid[j][i])
		}

		d += maxRow + maxCol
	}

	return d
}
