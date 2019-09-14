package hashTable

import "strings"

// 290
func wordPattern(pattern string, str string) bool {
	pm := map[string]int{}
	sm := map[string]int{}

	strs := strings.Split(str, " ")

	if len(strs) != len(pattern) {
		return false
	}

	for i:=0; i < len(pattern); i++ {
		pm[string(pattern[i])] += i + 1
		sm[strs[i]] += i + 1

		if pm[string(pattern[i])] != sm[strs[i]] {
			return false
		}
	}

	return true
}
