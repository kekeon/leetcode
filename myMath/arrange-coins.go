package myMath

import "math"

//441
/**
等差数列求和公式 S(n) = nA1 + (n-1)nd/2
解方程，求根公式，只能为正数
n = Sqrt(8S(n)+1)/2 - 0.5
再向下取整
 */
func arrangeCoins(n int) int {

	return int(math.Sqrt(float64(n*8) + 1) / 2 - 0.5)
}