package hash_table

func distributeCandies(candies []int) int {
	l := len(candies)

	m := map[int]int{}
	countKey := 0

	for _, v := range candies {
		m[v]++
		if m[v] == 1 {
			countKey++
		}
	}

	if l/2 >= countKey {
		return countKey
	}

	return l / 2
}
