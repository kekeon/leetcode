package myMath

import "strconv"

// 1154
func dayOfYear(date string) int {

	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// 2019-01-01
	// 月份
	y, _ := strconv.Atoi(date[0:4])
	m, _ := strconv.Atoi(date[5:7])
	d, _ := strconv.Atoi(date[8:10])

	num := 0

	if m == 1 {
		return d
	}

	// 判断世纪年， 能被100 整出的数为世纪年
	if y % 100 == 0 && y % 400 == 0 {
		days[1] = 29
	} else if y % 4 == 0 && y % 100 != 0 {
		days[1] = 29
	}

	for i := 0; i < m - 1; i++ {
		num += days[i]
	}



	return num + d

}
