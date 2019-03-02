package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {

	sum := 0
	s := []int{}
	var c int
	for _, v := range ops {

		switch v {

		case "+":
			c = s[len(s)-1] + s[len(s)-2]
			s = append(s, c)
			sum += c
			break
		case "D":

			c = s[len(s)-1] * 2
			s = append(s, c)
			sum += c
			break
		case "C":
			c = s[len(s)-1]
			s = s[0 : len(s)-1]
			sum -= c
			break
		default:
			c, _ = strconv.Atoi(v)
			s = append(s, c)
			sum += c
		}
	}

	return sum
}

func main() {
	num := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	v := calPoints(num)

	fmt.Println(v)
}
