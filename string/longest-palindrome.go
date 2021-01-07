package main

import "fmt"

/*
与第3题类似
思考点：最长子串，且是回文字符串，先找最长字串从最长的开始找
 */

// 返回值大于 0 则是回文字符串
func palindromeStr(s string) int {

	start := 0
	last := len(s) - 1

	for last-start > 0 {
		if s[start] == s[last] {
			start ++
			last --
		} else {
			return 0
		}
	}

	return len(s)
}

// w 窗口长度

func longestPalindrome(s string, start, maxL int, maxStr string) string {

	l := len(s)
	if l <= 1 {
		return s
	}

	sv := palindromeStr(s)

	if sv == l {
		return s
	}



	for i := start + 1; i < l; i++ {
		s1 := s[start:i]
		v1 := palindromeStr(s1)
		fmt.Println(s1, ":=", v1,  "maxL :=" , maxL, maxStr)

		if v1 > maxL {
			maxL = v1
			maxStr = s1
		}
	}

	start++
	return longestPalindrome(s[start:], 0, maxL, maxStr)
}

func main() {
	s := "aacabdkacaa"
	v := longestPalindrome(s, 0, 0, "")
	fmt.Println("===========")
	fmt.Println(v)
	// fmt.Println(palindromeStr(s))
}
