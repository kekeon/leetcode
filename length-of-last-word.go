package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if s == "" || s == " " {
		return 0
	}

	ss := strings.Split(s, "")

	n := 0

	for i:=0; i<len(ss);i++ {

		if i >= 1 {
			if ss[i]!= " " && ss[i-1] == " " {

				n = 1

			}else if ss[i] != " "{

				n++
			}
		}else if s[i]!=' ' {
			n++
		}

	}

	return n

}

func main() {
	s := lengthOfLastWord("  A")
	fmt.Println(s)
}
