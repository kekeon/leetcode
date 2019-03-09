package main

import (
	"fmt"
)

func reverseOnlyLetters(s string) string {

	m := map[string]bool{
		"A": true,
		"B": true,
		"C": true,
		"D": true,
		"E": true,
		"F": true,
		"G": true,
		"H": true,
		"I": true,
		"J": true,
		"K": true,
		"L": true,
		"M": true,
		"N": true,
		"O": true,
		"P": true,
		"Q": true,
		"R": true,
		"S": true,
		"T": true,
		"U": true,
		"V": true,
		"W": true,
		"X": true,
		"Y": true,
		"Z": true,
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"f": true,
		"g": true,
		"h": true,
		"i": true,
		"j": true,
		"k": true,
		"l": true,
		"m": true,
		"n": true,
		"o": true,
		"p": true,
		"q": true,
		"r": true,
		"s": true,
		"t": true,
		"u": true,
		"v": true,
		"w": true,
		"x": true,
		"y": true,
		"z": true,
	}

	start := 0

	l := len(s)

	last := l - 1

	word := ""
	str := ""

	for i := 0; i < l; i++ {

		if !m[string(s[i])] {

			word = ""

			for ; last >= 0; last-- {

				if m[string(s[last])] && len(word) != i-start {
					word += string(s[last])

				} else if len(word) == i-start {
					str += word
					break
				}

				if last == 0 {
					str += word
				}
			}
			str += string(s[i])
			start = i + 1
		} else if i+1 == l {
			word = ""
			for ; last >= 0; last-- {

				if m[string(s[last])] {
					word += string(s[last])
				}

				if last == 0 {
					str += word
				}
			}
		}

	}

	return str
}

func main() {
	v := reverseOnlyLetters("Test1ng-Leet=code-Q!")
	fmt.Println(v)
	//"j-Ih-gfE-dCba"
}
