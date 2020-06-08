package array

import "fmt"

func numIslands(grid [][]byte) int {

	num := 0

	g := &grid

	for x, level := range grid {
		for y, v := range level {
			if v == '0' {
				continue
			}
			num += sink(g, x, y)
		}
	}

	return num
}

func sink(g *[][]byte, i, j int) int {

	dx := []int{-1, -1, 0, 0}
	dy := []int{0, 0, -1, 1}

	fmt.Println((*g))
	if (*g)[i][j] == '0' {
		return 0
	}

	(*g)[i][j] = '0'

	for i :=0; i < len(dx); i++ {
		x := dx[i] + i
		y := dy[i] + j

		if x > 0 && x < len(*g) && y > 0 && y < len((*g)[x]) {
			if (*g)[i][j] == '0' {
				continue
			}

			sink(g, x, y)
		}
	}

	return 1
}
