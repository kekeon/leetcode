package hash_table

// 748
func shortestCompletingWord(licensePlate string, words []string) string {

	a := 97
	z := 122
	A := 65
	Z := 90

	licenseMap := map[string]int{}

	for _, v := range licensePlate {
		if int(v) >= a && int(v) <= z {
			licenseMap[string(v)]++
		}else if int(v) >= A && int(v) <= Z {
			licenseMap[string(v + 32)]++
		}
	}

	// 我们保证一定存在一个最短完整词。当有多个单词都符合最短完整词的匹配条件时取单词列表中最靠前的一个。
	l := 16 // 第一个完次出现的长度
	index := -1
	for i, word := range words {
		m := map[string]int{}

		for k, v := range licenseMap {
			m[string(k)] = v
		}

		for _, w := range word {
			m[string(w)]--
		}


		tag := true
		for _, v := range m {
			if v > 0 {
				tag = false
				break
			}
		}

		if tag && len(word) < l{
			l = len(word)
			index = i
		}
	}


	if index >= 0 {
		return words[index]
	}

 	return  ""
}
