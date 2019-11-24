package myMath

func minCostToMoveChips(chips []int) int {
	odd := 0
	even := 0

	for _, v := range chips {
		if v&1 == 0 {
			even++
		} else {
			odd ++
		}
	}

	if odd > even {
		return even
	}

	return odd
}
