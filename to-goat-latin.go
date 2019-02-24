package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(S string) string {

	mp := map[string]bool{
		"a":true,
		"e":true,
		"i":true,
		"o":true,
		"u":true,
		"A":true,
		"E":true,
		"I":true,
		"O":true,
		"U":true,
	}
	word := strings.Split(S, " ")
	last := ""

	for i,v := range word {

		last = strings.Repeat("a", i + 1)

		if mp[string(v[0:1])] {
			word[i] = word[i] + "ma" + last
		}else {
			word[i] = word[i][1:] + word[i][0:1] + "ma"+ last
		}

	}


	return strings.Join(word, " ")
}

func main() {
	v := toGoatLatin("I speak Goat Latin")
	fmt.Println(v)
}
