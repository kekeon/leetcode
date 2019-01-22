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

		rs = ""
		for j := 1; len(s) > 0; j++ {

			fmt.Println("--------------")
			fmt.Println(rs)
			fmt.Println(j)
			fmt.Println(s)
			fmt.Println(i)
			fmt.Println(p.v)
			fmt.Println("------end--------")
			cs := string(s[j])

			if p.v != cs {
				if len(s) == j+1  {

					num := j+1 - p.num
					rs = rs + strconv.Itoa(num) + string(s[j])
					j = 0
					s = ""
				}else {
					num := j - p.num
					s = s[p.num : j+1]
					rs = rs + strconv.Itoa(num) + p.v
					j = 0
				}
				p.v = cs
				p.num = j
			}else if len(s) == j+1 {
				m[i] = strconv.Itoa(j+1) + string(s[0])
				rs = rs + strconv.Itoa(j+1) + string(s[0])
				break
			}
		}

	}

	fmt.Println(m)
	return rs
}

func main() {
	v := number(4)
	fmt.Println("===========")
	fmt.Println(v)
}
