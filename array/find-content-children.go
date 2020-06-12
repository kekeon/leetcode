package array

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	gn := 0
	sn := 0
	count := 0

	for gn < len(g) && sn < len(s) {
		if g[gn] <= s[sn] {
			count ++
			gn ++
			sn ++
		} else {
			sn ++
		}
	}

	return count
}
