package array

import "strconv"

func letterCombinations(digits string) []string {

	res := []string{}

	if len(digits) == 0 {
		return res
	}

	zmap := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}


	search("", &res, digits, 0, zmap)
	return res
}

func search(s string, result *[]string, digits string, level int, m map[int]string) {

	if len(digits) <= level {
		*result = append(*result, s)
		return
	}

	i, _ := strconv.Atoi(digits[level: level + 1])
	str := m[i]

	level += 1
	for _, v := range str {
		search(s + string(v), result, digits, level, m)
	}
}
