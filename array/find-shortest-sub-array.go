package array

import (
	"math"
)

func findShortestSubArray(nums []int) int {


	if len(nums) == 1 {
		return 1
	}

	m := map[int]map[string]int{}
	maxCount := 0

	for i, v := range nums {

		if m[v] != nil {
			m[v]["count"]++
			m[v]["lastIndex"] = i
		}else {
			m[v] = map[string]int{
				"count" : 1,
				"lastIndex" : i,
				"startIndex" : i,
			}
		}

		if m[v]["count"] > maxCount {
			maxCount = m[v]["count"]
		}
	}

	minL := math.MaxInt64
	for k := range m {
		if m[k]["count"] == maxCount && m[k]["lastIndex"] - m[k]["startIndex"] + 1 < minL{
			minL =  m[k]["lastIndex"] - m[k]["startIndex"] + 1
		}
	}

	return minL
}
