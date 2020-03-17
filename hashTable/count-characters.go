package hashTable

// 1160
func countCharacters(words []string, chars string) int {

	sm := map[string]int{}
	for _, v := range chars {
		sm[string(v)]++
	}

	l := 0

	for _, word := range words {
		m := map[string]int{}
		tag := true
		for _, w := range word {
			m[string(w)] ++
			if sm[string(w)] - m[string(w)] < 0 {
				tag = false
				break
			}
		}

		if tag {
			l += len(word)
		}
	}

	return l
}
