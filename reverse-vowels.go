package main

import (
	"fmt"
	"strings"
)

func boolVowel(s string) bool {

	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "A" || s == "E" || s == "I" || s == "O" || s == "U" {
		return true
	}

	/*ss := "aeiouAEIOU"
	b := strings.Contains(ss, s)*/

	return false
 }

func reverseVowels(s string) string {

	sl := len(s)

	if sl <=1 {
		return s
	}

	list:= strings.Split(s,"")

	i := 0
	j := sl -1
	for i < j {

		fmt.Println(j)
		fmt.Println(i)

		if boolVowel(string(list[i])) &&  boolVowel(string(list[j])) {
			list[i], list[j] = list[j], list[i]
			i++
			j--
		}else if !boolVowel(string(list[i])) {
			i++
		}else if !boolVowel(string(list[j])) {
			j--
		}
	}

	return strings.Join(list,"")
}

func main() {
	v := reverseVowels("leetcode")
	fmt.Println(v)
}
