package myMath

import (
	"strconv"
)

func mySqrt(x int) int {

	str := strconv.Itoa(x)

	step := len(str)
	v := myPow(10, step/2)
	s := v * v
	for {

		if s <= x {
			if s == x || s+(v<<1)+1 > x {
				return v
			}

			step = len(strconv.Itoa(x - (v * v)))

			v = v + step

		} else {
			v = v / 2
		}

		s = v * v
	}

	return v

}
