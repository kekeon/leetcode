package main

import (
	"fmt"
	"strconv"
)

type sp struct {
	num int
	v   string
}

func number(n int) string {

	m := make(map[int]string)

	m[1] = "1"
	m[2] = "11"

	rs := ""

	if n <= 2 {
		return m[n]
	}

	for i := 3; i <= n; i++ {

		s := m[i-1]

		p := sp{
			num: 0,
			v:   "1",
		}

		for j := 0; len(s)-1 > 0; j++ {

			fmt.Println(j)
			cs := string(s[j])

			if p.v != cs || len(s) == j+1 {
				p.v = cs
				p.num = j - p.num + 1
				s = s[p.num : j+1]
				rs = rs + strconv.Itoa(p.num) + p.v
			}
		}

	}

	return rs
}

func main() {
	v := number(3)
	fmt.Println(v)
}
