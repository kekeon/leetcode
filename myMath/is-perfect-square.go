package myMath
// 367
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	n := num
	for n * n > num {
		n = ( n + num/n ) / 2
	}

	if n * n == num {
		return true
	}

	return false

}
