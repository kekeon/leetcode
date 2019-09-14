package hashTable

func numJewelsInStones(J string, S string) int {
	typeJ := map[string]bool{}

	count := 0

	for _, v := range J {
		typeJ[string(v)] = true
	}

	for _, v := range S {
		if typeJ[string(v)] {
			count++
		}
	}

	return count
}
