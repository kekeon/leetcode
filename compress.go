package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	l, num := 0, 1

	for i, v := range chars {
		if i+1 < len(chars) && v == chars[i+1] {
			num ++
		} else {
			chars[l] = v
			l ++
			if num > 1 {
				for _, num := range strconv.Itoa(num) {
					chars[l] = byte(num)
					l++
					fmt.Println(l)
				}
			}
			num = 1
		}
	}
	fmt.Println(chars)

	return l
}

func main() {

	v := compress([]byte{97, 99, 99, 99, 99, 99, 99, 99})
	fmt.Println(v)
}
