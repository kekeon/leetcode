package main

import "fmt"

// 滑动窗口
/**

startIndex： 起始位
maxLength： 当前窗口最大长度

 */
func lengthOfLongestSubstringMax(s string, startIndex, maxLength int) (i , maxL int){

	m := map[string]int{}
	maxL = 0
	sl := len(s)

	for i = startIndex; i < sl; i++ {
		sv := string(s[i])

		if m[sv] > 0 {
			if maxLength > maxL {
				maxL = maxLength
			}
			return lengthOfLongestSubstringMax(s, startIndex + 1, maxL)
		} else {
			m[sv] ++
			maxL ++
		}
	}

	if maxLength > maxL {
		maxL = maxLength
	}

	return i, maxL
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return  len(s)
	}

	_, l := lengthOfLongestSubstringMax(s, 0, 1)

	return l
}

func main() {

	l := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(l)

}
