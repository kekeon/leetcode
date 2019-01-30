package main

import (
	"strings"
	"math"
	"fmt"
)

func isPalindrome(s string) bool{

	l := len(s)

	if l == 0 {
		return false
	}

	ss := strings.ToUpper(s)
	list := []string{}

	for a:=0; a < l; a++ {
		if ss[a] <= 90 && ss[a] >= 65 || ss[a] <= 57 && ss[a] >= 48 {
			list = append(list, string(ss[a]))
		}
	}
	fmt.Println(list)

	ssl := len(list)

	var f float64
	f = 3/2
	fmt.Println(f)
	fmt.Println(3.9)
	fmt.Println("----------------====")
	num := int(math.Round(float64(ssl / 2)))

	lab := false

	if ssl % 2 == 0 {
		lab = true
	}

	fmt.Println("+++++")
	fmt.Println(num)
	for i:= 0; i < num; i++{

		fmt.Println(i)
		if !lab && i + 1 == num {
			fmt.Println("---------------")
			return true
		}

		if list[i] != list[ssl-i - 1] {
			fmt.Println("============")
			return false
		}
	}

	return true
}

func main() {
	v := isPalindrome("abb")
	fmt.Println(v)
}
