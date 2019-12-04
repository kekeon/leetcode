package myMath

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a := 1
	b := 2
	c := 0

	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return b

}
