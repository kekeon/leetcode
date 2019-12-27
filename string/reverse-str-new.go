package main

// s = "abcdefg", k = 2, bacdfeg

/*
 三步反转法
 bagfedc -> cdefgab

 X -> X^T
 Y -> Y^T
 (X^TY^T)^T=YX

*/

// 541. 反转字符串2
func reverseStrNew(s string, k int) string {

	l := len(s)
	sByte := []byte(s)
	// 余长
	c := l % (2 * k)
	n := (l - c) / (2 * k)

	// 反转正常大于2K // 0 ~ k , 2k ~ 2k + k
	for i := 0; i < n; i++ {
		reverseStringByte(i * k * 2 , i * k * 2 + k - 1, sByte)
	}

	// 大于等于k 反转
	if c >= k {
		reverseStringByte(l - c , l - c + k - 1, sByte)
	// 反转小于K 全部
	} else if c > 1 && c < k {
		reverseStringByte(l - c , l - 1, sByte)
	}


	return string(sByte)
}



func reverseStringByte(start, end int, s []byte) {

	c := s[start]
	for start < end {
		s[start] = s[end]
		s[end] = c
		start++
		end--
		c = s[start]
	}
}
// krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc
//"jlnnxsetgcpsbhsfymrkhfursyissjnsocgdhgfxtxrlvugsaphqzxllwebukgatzfybprfmmfithphccxfsogsgqsnvckjvnskk"
//"jlnnxsetgcpsbhsfymrkhfursyissjnsocgdhgfxtxrlvugsaphqzxllwebukgatzfybprfmmfithphccxfsogsgqsnvckjvnskk"
func main() {
	v := reverseStrNew("krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc", 20)
	println(v)
}
