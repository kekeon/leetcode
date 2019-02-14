package main

import (
	"fmt"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {

	for _, v := range ransomNote {
		sub := string(v)

		index := strings.Index(magazine, sub)

		if index != -1 {
			magazine = strings.Replace(magazine, sub, "-", 1)
		} else {
			return false
		}

	}

	return true
}

func main() {
	v := canConstruct("a", "b")
	fmt.Println(v)
}
