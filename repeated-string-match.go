package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(A string, B string) int {

	n := 1

	num := 0

	al := len(A)
	bl := len(B)

	if al > bl {
		num = 2
	} else {
		num = len(B) / len(A) * 3
	}

	str := A

	for {

		i := strings.Index(str, B)

		if i != -1 {
			return n
		}

		if n >= num {
			return -1
		}

		str += A
		n++
	}

}

func main() { //abcabcabc
	v := repeatedStringMatch("aaaaaaaaaaaaaaaaaaaaaab", "ba")
	fmt.Println(v)
}
