package bit_manipulation

func hammingWeight(num uint32) int {

	c := 0
	for num > 0 {

		if num & 1 == 1{
			c++
		}
		num >>=1
	}

	return c
}