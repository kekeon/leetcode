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
			v:   string(s[0]),
		}

		for j := 0; len(s) > 0; j++ {
			fmt.Println(s)
			cs := string(s[j])

			if p.v != cs {
				fmt.Println(i)
				if len(s) == j+1  {
					s = ""
					num := j - p.num
					rs = rs + strconv.Itoa(num) + p.v
					j = 0
				}else {
					num := j - p.num
					s = s[p.num : j+1]
					rs = rs + strconv.Itoa(num) + p.v
					j = 0
				}
				p.v = cs
				p.num = j
			}else if len(s) == j+1 {

				
				break
			}
		}

	}

	return rs
}

func main() {
	v := number(4)
	fmt.Println(v)
}
