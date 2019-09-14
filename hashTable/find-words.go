package hashTable

import (
	"strings"
)

func findWords(words []string) []string {

	res := []string{}

	for _, word := range words {
		if b, v := isInclude(word[0:1]); b {
			ok := true
			for _, s := range word {
				k := strings.ToLower(string(s))
				if !v[k] {
					ok = false
					break
				}
			}

			if ok {
				res = append(res, word)
			}
		}
	}

	return res
}

func isInclude(s string) (bool, map[string]bool) {
	colMap1 := map[string]bool{
		"q": true,
		"w": true,
		"e": true,
		"r": true,
		"t": true,
		"y": true,
		"u": true,
		"i": true,
		"o": true,
		"p": true,
	}

	colMap2 := map[string]bool{
		"a": true,
		"s": true,
		"d": true,
		"f": true,
		"g": true,
		"h": true,
		"j": true,
		"k": true,
		"l": true,
	}

	colMap3 := map[string]bool{
		"z": true,
		"x": true,
		"c": true,
		"v": true,
		"b": true,
		"n": true,
		"m": true,
	}

	k := strings.ToLower(s)

	if colMap1[k] {
		return true, colMap1
	}

	if colMap2[k] {
		return true, colMap2
	}

	if colMap3[k] {
		return true, colMap3
	}

	return false,map[string]bool{}
}
