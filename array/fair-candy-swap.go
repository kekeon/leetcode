package array

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0

	 mB := map[int]int{}

	for _, v := range A {
		sumA += v
	}

	for _, v := range B {
		sumB += v
		mB[v] = v
	}

	p := ( sumB - sumA ) / 2


	for _, v := range A {
		if ( mB[p+v] > 0) {
			return []int{v, mB[p+v]}
		}
	}

	return []int{}
}
