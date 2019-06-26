package array

func canThreePartsEqualSum(A []int) bool {
	sum := 0

	for _, v := range A {
		sum += v
	}

	if sum % 3 != 0 {
		return false
	}

	avg := sum / 3

	s := 0
	for _, v := range A {
		s += v

		if s == avg {
			s = 0
		}
	}


	if s == 0 {
		return true
	}

	return false
}