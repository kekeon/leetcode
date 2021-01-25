package myMath

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {

	res := []string{}

	restoreIpAddressesRecursion(s, 0, []string{}, &res)

	return res
}

func restoreIpAddressesRecursion(s string, start int, path []string, res *[]string) {


	// 终止条件,并满足结果
	if len(path) == 4 && start == len(s) {
		*res = append(*res, strings.Join(path, "."))
		return
	}

	// 终止条件不满足，
	if  (start < len(s) && len(path) == 4) ||
		(start == len(s) && len(path) < 4) {
		return
	}

	for a := 1; a <=3 ; a++ {
		// 加上要切的长度就越界，不能切这个长度
		if start + a > len(s){
			return
		}

		// 后三位不能以0开头
		if a !=1 && s[start] == '0' {
			return
		}

		// 大于255
		str := s[start: start + a]
		sv, _ := strconv.Atoi(str)
		if a == 3 && sv > 255 {
			return
		}

		path = append(path, str)
		restoreIpAddressesRecursion(s, start + a, path, res)
		path = path[0: len(path) - 1]
	}
}
