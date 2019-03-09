package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {

	if len(s) == 1 {
		return false
	}

	l := len(s)

	for i:= 1; i<=l / 2; i++ {
		sub := s[0:i]

		if l%len(sub) == 0 && strings.Repeat(sub,l / len(sub)) == s {
			return true
		}
	}

	return false
}

func main() {
	v := repeatedSubstringPattern("abaabaaba")
	fmt.Println(v)
}
