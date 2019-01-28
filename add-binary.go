package main

import (
	"fmt"
	"strings"
	"strconv"
)

func addBinary(a string, b string) string {

	if a == "0" {
		return b
	}

	if b == "0" {
		return a
	}

	al := len(a)
	bl := len(b)

	l := 0

	if al >= bl {
		l = al
	} else {
		l = bl
	}

	aList := strings.Split(a, "")
	bList := strings.Split(b, "")

	list := []string{}

	lab := false

	for i:=0; i<=l; i++ {

		aNumber := 0
		bNumber := 0
		if i<len(aList) {
			aNumber,_ = strconv.Atoi(aList[al - 1 - i])
		}

		if i<len(bList) {
			bNumber,_ = strconv.Atoi(bList[bl - 1 - i])
		}

		if i== l {
			if lab {
				list = append(list,"1")
			}
			break
		}

		sum := aNumber + bNumber

		if sum >= 2 {

			if lab {
				list = append(list,"1")
			}else {
				list = append(list,"0")
				lab = true
			}

		}else {

			if lab {
				sum = sum + 1
			}

			if sum >= 2 {
				list = append(list,"0")
				lab = true
			}else {
				s := strconv.Itoa(sum)
				list = append(list,s)
				lab = false
			}

		}

	}

	str := ""

	for j:=len(list) -1; j>=0; j-- {
		str += list[j]
	}


	return str
}

func main() {
	fmt.Println(addBinary("100", "110010")) // 100
	fmt.Println(addBinary("1010", "1011")) // 10101
}
