package stack

import "fmt"

func backspaceCompare(S string, T string) bool {

	ss := ""

	tt := ""

	l := 0

	for _, v := range S {

		if string(v) != "#" {
			ss += string(v)
		} else {
			l = len(ss)
			if l > 0 {
				ss = ss[:l-1]
			}
		}
	}

	for _, v := range T {

		if string(v) != "#" {
			tt += string(v)
		} else {
			l = len(tt)
			if l > 0 {
				tt = tt[:l-1]
			}
		}
	}

	if tt == ss {
		return true
	}else {
		return false
	}
}

func testBackspaceCompare() {

	v := backspaceCompare("a#c", "c")

	fmt.Println(v)
}
