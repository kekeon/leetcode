package myMath

func countPrimes(n int) int {
	c := 0
	arr := make([]bool, n)

	for i := 2; i < n; i++ {
		if (!arr[i-1]) {
			c++
			for j := i * i; j <= n; j += i {
				arr[j-1] = true
			}
		}
	}

	return c
}
