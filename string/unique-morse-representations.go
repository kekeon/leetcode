package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {

	m := map[string]string{
		"a" : ".-",
		"b" : "-...",
		"c" : "-.-.",
		"d" : "-..",
		"e" : ".",
		"f" : "..-.",
		"g" : "--.",
		"h" : "....",
		"i" : "..",
		"j" : ".---",
		"k" : "-.-",
		"l" : ".-..",
		"m" : "--",
		"n" : "-.",
		"o" : "---",
		"p" : ".--.",
		"q" : "--.-",
		"r" : ".-.",
		"s" : "...",
		"t" : "-",
		"u" : "..-",
		"v" : "...-",
		"w" : ".--",
		"x" : "-..-",
		"y" : "-.--",
		"z" : "--..",
	}

	mose := ""

	toMose := map[string]int{}

	count := 0
	for _, v := range words {
		mose = ""

		for _,w := range string(v) {
			mose += m[string(w)]
		}
		toMose[mose]++

		if toMose[mose] == 1 {
			count ++
		}

	}

	return count
}

func main() {
	m := []string{"gin", "zen", "gig", "msg"}
	v := uniqueMorseRepresentations(m)
	fmt.Println(v)
}
