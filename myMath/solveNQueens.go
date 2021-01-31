package myMath

func solveNQueens(n int) [][]string {

	result := [][]string{}
	path := make([]int, n)
	for i := 0; i < n; i++ {
		path[i] = -1
	}

	solveNQueensRecursion(n, map[int]bool{}, map[int]bool{}, map[int]bool{}, 0, path , &result)

	return result
}

func solveNQueensRecursion(n int, rowMap, diffXY, sumXY map[int]bool, col int, path []int, res *[][]string) {

	// 终止条件
	if col >= n {
		board := generateBoard(path, n)
		*res = append(*res, board)
		return
	}

	for i := 0; i < n; i++ {
		if rowMap[i] || sumXY[col + i] || diffXY[col - i] {
			continue
		}

		rowMap[i] = true
		sumXY[col + i] = true
		diffXY[col - i] = true

		path[col] = i
		solveNQueensRecursion(n, rowMap, diffXY, sumXY, col+1, path, res)
		rowMap[i] = false
		sumXY[col + i] = false
		diffXY[col - i] = false
		path[col] = -1
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
