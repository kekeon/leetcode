package myMath

import "math"
// 直线斜率相等
func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates)

	if l <= 2 {
		return  true
	}

	a := coordinates[0]
	b := coordinates[1]
	k := math.Abs((float64(a[1] - b[1])) / (float64(a[0] - b[0])))
	var kk float64
	for i := 2; i < l; i++ {
		a = coordinates[i]
		b = coordinates[i-1]
		kk =   math.Abs((float64(a[1] - b[1])) / (float64(a[0] - b[0])))

		if kk != k {
			return false
		}
	}

	return true
}
