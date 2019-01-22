package main

import (
	"fmt"
	"strconv"
)

type sp struct {
	num int
	v   string
}

func countAndSay(n int) string {

	m := make(map[int]string)

	m[1] = "1"
	m[2] = "11"

	rs := ""

	if n <= 2 {
		return m[n]
	}

	for i := 3; i <= n ; i++ {
		s := m[i-1]

		p := sp{
			num: 0,
			v:   string(s[0]),
		}

		rs = ""
		for j := 1; len(s) > 0; j++ {
			cs := string(s[j])

			if p.v != cs {
				if len(s) == j+1  {
					if p.v == string(s[j]) {
						num := j+1 - p.num
						rs = rs + strconv.Itoa(num) + string(s[j])
						j = 0
						s = ""
					}else {
						num := j - p.num
						rs = rs + strconv.Itoa(num) + p.v
						rs = rs + "1" + string(s[j])
						j = 0
						s = ""
					}
					m[i] = rs
					break
				}else {
					num := j - p.num
					s = s[j:]
					rs = rs + strconv.Itoa(num) + p.v
					j = 0
				}
				p.v = cs
				p.num = j
			}else if len(s) == j+1 {
				rs = rs + strconv.Itoa(j+1) + string(s[0])
				m[i] = rs
				break
			}
		}

	}

	return rs
}

func main() {
	v := countAndSay(30)
	fmt.Println(v)
}
