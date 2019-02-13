package main

import (
	"strings"
	"math"
	"fmt"
)

func isPalindrome(s string) bool{

	l := len(s)

	if l == 0 {
		return true
	}

	ss := strings.ToUpper(s)
	list := []string{}

	for a:=0; a < l; a++ {
		if ss[a] <= 90 && ss[a] >= 65 || ss[a] <= 57 && ss[a] >= 48 {
			list = append(list, string(ss[a]))
		}
	}
	fmt.Println(list)

	ssl := float64(len(list))
	ssl_int := len(list)


	num := int(math.Round(ssl / 2))

	lab := false

	if ssl_int % 2 == 0 {
		lab = true
	}

	for i:= 0; i < num; i++{

		if !lab && i + 1 == num {
			return true
		}

		if list[i] != list[ssl_int-i - 1] {
			return false
		}
	}

	return true
}

func main() {
	v := isPalindrome("")
	fmt.Println(v)
}
