package hash_table

func findAnagrams(s string, p string) []int {

	subL := len(p)
	res := []int{}
	subMap := map[string]int{}
	for _, k := range p{
		subMap[string(k)]++
	}


	for i := 0; i <= len(s) - subL; i++ {
		subS := s[i:i+subL]
		isOk := true
		subM := map[string]int{}
		for _, subK := range subS{
			subM[string(subK)]++
		}

		for k, v := range subMap {
			if subM[string(k)] != v {
				isOk = false
				break
			}
		}

		if isOk {
			res = append(res, i)
		}
	}

	return res

}
