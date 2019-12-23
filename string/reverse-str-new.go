package main


// s = "abcdefg", k = 2, bacdefg

/*
 三步反转法
 bagfedc -> cdefgab

 X -> X^T
 Y -> Y^T
 (X^TY^T)^T=YX

*/

// 541. 反转字符串2， 只需要反转前面后拼接后面

func reverseStrNew(s string, k int) string {
	return reverseStringNew(0, k -1 , s)
}

func reverseStringNew(start, end int, s string) string {

	s2 := []byte(s)
	c := s2[start]
	for start < end {
		s2[start] = s[end]
		s2[end] = c
		start++
		end--
		c = s2[start]
	}

	return string(s2)
}

func main() {
	v := reverseStrNew("abcdefg", 2)
	println(v)
}
