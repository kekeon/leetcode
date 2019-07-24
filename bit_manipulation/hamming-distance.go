package bit_manipulation

func hammingDistance(x int, y int) int {

	n := x^y
	c := 0
	for n > 0 {
		c++
		n = n&(n-1)
	}

	return c

}