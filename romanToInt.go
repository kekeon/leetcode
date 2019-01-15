package main

import "fmt"

// 13. 罗马数字转整数
func romanToInt(s string) int {

	sum := 0

	if len(s) == 0 {
		return sum
	}

	m := make(map[string]int)

	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	ls := len(s)
	for i := 0; i < ls; i++ {

		a := m[string(s[i])]
		b := 0

		// 判断是否越界
		if (i+1 < ls) {
			b = m[string(s[i+1])]
		}

		// 前着是否小于后着
		if (a < b) {
			sum += b - a
			i++
		} else {
			sum += a
		}

	}

	return sum
}

func main() {
	v := romanToInt("MCMXCIV")
	fmt.Println(v)
}
