package hashTable

func islandPerimeter(grid [][]int) int {
	h := len(grid)

	w := len(grid[0])

	count := 0

	for i :=  0; i < h; i++ {
		for j := 0; j < w; j++ {

			if grid[i][j] == 1 {

				if j == 0 || grid[i][j-1] == 0 {
					count++
				}

				if i == 0 || grid[i-1][j] == 0 {
					count++
				}
			}

		}
	}

	return count *  2

}
