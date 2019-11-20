package myMath

func bitwiseComplement(N int) int {
	s :=  1
	for s < N {
		s = s << 1 | 1
	}

	return s - N
}