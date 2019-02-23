package main

import (
	"fmt"
	"strings"
	"regexp"
)

func mostCommonWord(paragraph string, banned []string) string {

	paragraph = strings.ToLower(paragraph)

	re := regexp.MustCompile("((!|,|\\.|\\?|;|')\\s)|(!|,|\\.|\\?|;|')|\\s")

	word := re.Split(paragraph, -1)

	bannedWord := map[string]bool{}

	for _, v := range banned {
		bannedWord[string(v)] = true
	}

	mp := map[string]int{}

	resIndex := 0

	count := 0

	k := ""
	for i, v := range word {

		k = string(v)
		if bannedWord[k] {

			mp[k] = 0

		} else {
			mp[k] ++

			if mp[k] > count {
				count = mp[k]
				resIndex = i
			}

		}
	}

	return word[resIndex]

}

func main() {

	paragraph := "a, a, a, a, b,b,b,c, c"
	banned := []string{"a"}
	v := mostCommonWord(paragraph, banned)

	fmt.Println(v)
}
