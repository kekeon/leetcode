package main

import (
	"fmt"
	"strings"
)

// 20. 有效的括号
func bracketValid(s string) bool {

	isValid := false
	if len(s)%2 != 0 {
		return isValid
	}
	m := make(map[string]string)

	m["{"] = "}"
	m["["] = "]"
	m["("] = ")"

	i := 0

	for {
		l := len(s)

		if l-1 <= i {
			break
		}

		if m[string(s[i])] == string(s[i+1]) {
			sArr := strings.Split(s, "")

			sArr[i] = ""
			sArr[i+1] = ""
			s = strings.Join(sArr, "")
			i = 0
		} else {
			i++
		}
	}

	if len(s) != 0 {
		return false
	} else {
		return true

	}
}

func main() {
	v := bracketValid("[()](()(})")
	fmt.Println(v)
}
