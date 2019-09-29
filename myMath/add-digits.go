package myMath

func addDigits(num int) int {
	if num != 0 && num % 9 == 0 {
		return 9
	}

	return  num % 9
}