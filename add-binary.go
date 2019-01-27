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

	if al <= bl {
		l = al
	} else {
		l = bl
	}

	aList := strings.Split(a, "")
	bList := strings.Split(b, "")

	list := []string{}

	label := false  // 是否有进位
	for i := l-1; i >=0; i-- {
		aNumber,_ :=  strconv.Atoi(aList[i])
		bNumber,_ :=  strconv.Atoi(bList[i])

		n := aNumber + bNumber

		if n > 1 {
			if label {
				list = append(list, "1")
				label = false
			}else {
				list = append(list, "0")
				label = true
			}
		}else {
			s := strconv.Itoa(n)

			if !label {
				list = append(list, s)
			}else {
				n = n + 1

				if n > 1 {
					label = true
					list = append(list, "0")
				}else {
					s := strconv.Itoa(n)
					list = append(list, s)
				}

			}

		}
	}
	ss := ""
	if al > bl {
		ss = a[0:bl]
	} else if bl > al {
		ss = b[0:al]
	}

	sl := len(ss)
	slLast := 0
	if sl > 0 {
		slLast,_ = strconv.Atoi(ss[sl-1:sl]);
	}


	if label {

		if sl == 0 {
			list = append(list, "1")
		}else if slLast == 0{
			list = append(list, "1")
			for _, v := range ss[0:sl] {
				list = append(list, string(v))
			}
		}else {

			label = false
			for _, v := range ss {
				vNumber,_ := strconv.Atoi(string(v))

				n := vNumber +1


				if n > 1 {
					if label {
						list = append(list, "1")
						label = false
					}else {
						list = append(list, "0")
						label = true
					}
				}else {
					s := strconv.Itoa(n)

					if !label {
						list = append(list, s)
					}else {
						n = n + 1
						if n > 1 {
							label = true
							list = append(list, "0")
						}else {
							s := strconv.Itoa(n)
							list = append(list, s)
						}
					}

				}

			}

		}
	}else if sl > 0{
		for _, v := range ss {
			list = append(list, string(v))
		}
	}

	fmt.Println(list)
	return " "
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
