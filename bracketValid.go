package main

import (
	"strings"
	"fmt"
)

// 20. 有效的括号
func bracketValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}
	m := make(map[string]string)

	m["{"] = "}"
	m["["] = "]"
	m["("] = ")"

	i := 0
	isValid := false
	for {
		if i >= len(s) {

			if 0 == len(s) {
				isValid = true
			}

			break
		} else {
			// 判断 相邻一个是否，满足
			a := string(s[i])

			if i + 1 < len(s) && m[a] == string(s[i+1]) {

				var cs string
				if i > 0 {
					cs = s[0:i]
					s = s[i:]
				}

				s = strings.Replace(s, a, "", 1)
				s = strings.Replace(s, string(s[i]), "", 1)
				s = cs + s
				i = 0
				// 判断成对字符串
			} else if  i + 3 < len(s) && m[a] == string(s[i+3]) && string(s[i+1]) == string(s[i+2]) {
				s = strings.Replace(s, a, "", 1)
				s = strings.Replace(s, string(s[i+2]), "", 1)
				i = 0
			} else {
				i++
			}


			fmt.Println(s)
			fmt.Println(i)
		}
	}

	return isValid
}

func main() {
	fmt.Println("===")
	//v := "123"[0:1]

	v := bracketValid("[()](())")
	fmt.Println(v)
}
