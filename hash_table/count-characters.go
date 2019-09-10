package hash_table

// 1160
func countCharacters(words []string, chars string) int {
	count := 0
	cMap := map[string]int{}

	for _, s := range chars {
		cMap[string(s)]++
	}

	str, k := "", ""
	for _, w := range words {
		str = ""
		for _, v := range w {
			k = string(v)
			if cMap[k] > 0 {
				cMap[k] --
				str += k
			}
		}

		// 满足就累加
		if str == w {
			count += len(str)
		}

		// 还原
		for _, v := range str {
			k = string(v)
			cMap[k] ++
		}

	}



	return count
}
