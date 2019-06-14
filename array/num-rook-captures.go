package array

func numRookCaptures(board [][]byte) int {

	Rx, Ry, count := 0, 0, 0

	for x := 0; x < 8; x ++ {
		for y := 0; y < 8; y ++ {
			if board[x][y] == 'R' {
				Rx, Ry = x, y
				goto RIGHT
			}
		}
	}

	RIGHT:
		for r := Rx + 1; r < 8; r++ {
			switch board[r][Ry] {
			case 'p':
				count++
				break RIGHT
			case 'B':
				break RIGHT

			}
		}

	LEFT:
		for l := Rx - 1; l >= 0; l-- {
			switch board[l][Ry] {
			case 'p':
				count++
				break LEFT
			case 'B':
				break LEFT
			}
		}
	UP:
		for u := Ry - 1; u >= 0; u-- {
			switch board[Rx][u] {
			case 'p':
				count++
				break UP
			case 'B':
				break UP
			}
		}

	DOWN:
		for d := Ry + 1; d < 8; d++ {
			switch board[Rx][d] {
			case 'p':
				count++
				break DOWN
			case 'B':
				break DOWN
			}
		}
	return count
}
