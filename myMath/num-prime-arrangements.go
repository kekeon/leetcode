package myMath

func numPrimeArrangements(n int) int {

	if n < 2 {
		return 1
	}

	a := countP(n)
	b := n - a
	return arrangement(a) * arrangement(b) % 1000000007
}

func arrangement(n int) int {
	c := 1
	for i := 2; i <= n; i++ {
		c *= i
		c %= 1000000007
	}

	return c % 1000000007
}

func countP(n int) int {
	c := 0
	arr := make([]bool, n)

	for i := 2; i <= n; i++ {
		if !arr[i-1] {
			c++

			for j := i * i; j <= n; j += i {
				arr[j-1] = true
			}
		}
	}

	return c
}
