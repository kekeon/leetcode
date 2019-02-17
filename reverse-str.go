package main

import (
	"fmt"
	"math"
)

func reverseStr(s string, k int) string {

	sl := len(s)
	if sl < k {
		return reverseString(s)
	}

	n := int(math.Floor(float64(sl / (2 * k))))

	newStr := ""
	sub := ""
	rs := ""
	start, end := 0,0
	for i := 0; i < n; i ++ {
		start, end = i*2*k, (i+1)*2*k
		sub = s[start:end]
		rs = sub[0:k]

		newStr += reverseString(rs)  + sub[k:]
	}

	y := sl - end

	if y != 0 {
		if y < k {
			sub = s[end:]
			newStr += reverseString(sub)
		}else {
			sub = s[end:]
			rs := sub[0:k]
			newStr += reverseString(rs)  + sub[k:]
		}
	}

	return newStr
}

func reverseString(s string) string {
	runes := []rune(s)
	for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
		runes[first], runes[last] = runes[last], runes[first]
	}

	return string(runes)
}

func main() {
	v := reverseStr("abcdefg", 2)
	fmt.Println(v)
}
