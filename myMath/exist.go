package myMath

func exist(board [][]byte, word string) bool {

	row := len(board)
	col := len(board[0])
	mark := [][]int{}

	for i := 0; i < row; i++ {
		n := []int{}
		for j := 0; j < col; j++ {
			n = append(n, 0)
		}
		mark = append(mark, n)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				mark[i][j] = 1
				if existRecursion(board, j, i, word[1:], mark) {
					return true
				} else {
					mark[i][j] = 0
				}
			}
		}
	}

	return false
}

func existRecursion(board [][]byte, colIndex, rowIndex int, word string, mark [][]int) bool {
	directs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	if len(word) == 0 {
		return true
	}

	for _, v := range directs {
		i := colIndex + v[1]
		j := rowIndex + v[0]

		if j >= 0 && j < len(board) && i >= 0 && i < len(board[0]) && board[j][i] == word[0] {
			if  mark[j][i] == 1{
				continue
			}

			mark[j][i] = 1

			if existRecursion(board, i, j, word[1:], mark) {
				return true
			} else {
				mark[j][i] = 0
			}
		}

	}

	return false
}
