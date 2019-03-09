package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	list := []string{}
	cl := strings.Split(s, " ")
	for _, v := range cl {
		  list = append(list,ReverseString(string(v)))
	}

	return strings.Join(list, " ")
}

func ReverseString(s string) string {
	runes := []rune(s)
	for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
		runes[first], runes[last] = runes[last], runes[first]
	}

	return string(runes)
}



func main() {
	v := reverseWords("Let's take LeetCode contest")
	fmt.Println(v)
}
