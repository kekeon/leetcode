package bitManipulation

func countPrimeSetBits(L int, R int) int {

	sum := 0

	for i := L; i <= R; i++ {
		count := 0
		s := i

		for s != 0 {
			s &= s - 1
			count ++
		}

		if boolPrime(count) {
			sum ++
		}
	}

	return sum
}

func boolPrime(n int) bool {

	if n <= 3 {
		return n >= 2
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i < n; i += 6 {
		if n%2 == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
