package main

import (
	"fmt"
	"strconv"
)

func rotatedDigits(N int) int {

	m := map[string]bool{
		"0":true,
		"1":true,
		"2":true,
		"5":true,
		"6":true,
		"8":true,
		"9":true,
	}

	c := map[string]bool{
		"0":true,
		"1":true,
		"8":true,
	}

	count := 0

	k := ""

	isM := false

	isC := true

	for i := 1; i <= N; i++ {
		k = strconv.Itoa(i)

		if i < 10 && m[k] && !c[k] {
			count ++
		}

		if i > 11 {

			isM = true

			isC = true

			for _,v := range k {


				if !m[string(v)] {
					isM = false
					break
				}

				if !c[string(v)] && isC {
					isC = false
				}

			}

			if isM && !isC{
				count++
			}

		}
	}

	return count

}

func main() {
	v := rotatedDigits(20)
	fmt.Println(v)
}
