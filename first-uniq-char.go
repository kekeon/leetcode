package main

import (
	"fmt"
	"strings"
)

func firstUniqChar(s string) int {

	for i,v := range s {

		str := s[i+1:]
		sub := string(v)
		index := strings.Index(str, sub)

		if index != -1 {
			continue
		}
		str = s[:i]
		sub = string(v)
		index = strings.Index(str, sub)

		if index == -1 {
			return i
		}
	}

	return -1
}


func main() {
	v := firstUniqChar("aa")
	fmt.Println(v)
}
