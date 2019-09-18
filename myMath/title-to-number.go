package myMath

// 171
func titleToNumber(s string) int {
	m := map[string]int{
		"A":1,
		"B":2,
		"C":3,
		"D":4,
		"E":5,
		"F":6,
		"G":7,
		"H":8,
		"I":9,
		"J":10,
		"K":11,
		"L":12,
		"M":13,
		"N":14,
		"O":15,
		"P":16,
		"Q":17,
		"R":18,
		"S":19,
		"T":20,
		"U":21,
		"V":22,
		"W":23,
		"X":24,
		"Y":25,
		"Z":26,
	}

	l := len(s)

	if l == 1 {
		return m[s]
	}

	sum := 0

	for i := 1; i <=l; i++ {
		sum += m[s[i-1:i]] * myPow(26, l-i)
	}

	return sum

}

func titleToNumber2(s string) int {
	ans := 0
	for _,x := range s {
		t := int(x - 'A') + 1
		ans = ans*26 + t
	}
	return ans
}
