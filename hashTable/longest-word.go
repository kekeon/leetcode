package hashTable

import (
	"sort"
)

// 720
func longestWord(words []string) string {
	maxL := 0

	mapW := map[int][]string{}
	for _, v := range words {
		l := len(v)
		if l > maxL {
			maxL = l
		}

		if mapW[l] == nil {
			mapW[l] = []string{v}
		}else {
			mapW[l] = append(mapW[l], v)
		}
	}


	data := []string{}

	m := map[string]int{}
	upArr := mapW[1]
	for _,v1 := range upArr {
		m[v1] ++
	}
	for i := 2; i <= maxL; i++ {
		cArr := mapW[i]
		if len(data) > 0 {
			upArr = data
			data = []string{}
		} else if i > 2{
			break
		}

		for _,v2 := range cArr {
			k:= v2[0:i-1]
			if m[k] > 0 {
				data = append(data, v2)
				upArr = append(upArr, k)
				m[v2]++
			}
		}
	}

	if len(data) > 0 {
		upArr = data
	}

	if len(upArr) == 0 {
		return  ""
	}
	sort.Strings(upArr)
	return upArr[0]
}
