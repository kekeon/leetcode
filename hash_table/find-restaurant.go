package hash_table

func findRestaurant(list1 []string, list2 []string) []string {

	m := map[string]int{}

	sumIndex := 3000

	res := []string{}

	for i, v := range list1 {
		m[string(v)] = i + 1
	}

	for ii, v := range list2 {
		if m[string(v)] > 0 && (m[string(v)] + ii - 1) < sumIndex {
			sumIndex = m[string(v)] + ii - 1
		}
	}

	for ii, v := range list2 {
		if m[string(v)] > 0 &&  m[string(v)] + ii - 1 == sumIndex {
			res = append(res, v)
		}
	}

	return res
}
