package main

import "fmt"

// 28. 实现strStr()
func strStr(haystack string, needle string) int {

	b := 0
	n := len(needle)
	h := len(haystack)
	if n == 0 || (n == h && haystack == needle)  {
		return b
	}

	if n == h && haystack != needle {
		return -1
	}

	if h == 0 || n > h {
		return -1
	}

	i:=0
	fmt.Println(h - n)
	for  {
		fmt.Println(i)
		fmt.Println(haystack[i:i+n])

		if i <= h - n {

			if i == h-n && haystack[i:i+n] != needle {
				b = -1
				break
			}
			if haystack[i:i+n] == needle {
				b = i
				break
			}
		}else {
			b = -1
			break
		}

		i++
	}

	return b
}

func main() {
	v := strStr("aaaa", "bba")
	fmt.Println("==========")
	fmt.Println(v)
}
