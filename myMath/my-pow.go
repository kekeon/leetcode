package myMath

func myPow(a, p int) int {

	if p == 0 {
		return 1
	}

	r := a

	for i:= 1; i <= p; i++ {
		r *= a
	}

	return a
}
