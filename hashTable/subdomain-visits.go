package hashTable

import (
	"strings"
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
	strs := []string{}
	m := map[string]int{}


	for _, str := range cpdomains {
		// 拆解次数

		s := strings.Split(str, " ")
		num, _ := strconv.Atoi(s[0])
		domains := strings.Split(s[1], ".")

		// 每个域名地址包含一个或两个"."符号。
		if len(domains) == 2 {
			m[s[1]] += num
			m[domains[1]] += num
		} else if len(domains) == 3 {
			m[s[1]] += num
			m[domains[1]  + "." + domains[2]] += num
			m[domains[2]] += num
		}

	}

	for k, v := range m {
		strs = append(strs, strconv.Itoa(v) + " " + k)
	}


	return strs

}
