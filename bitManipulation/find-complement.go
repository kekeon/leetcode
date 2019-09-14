package bitManipulation

func findComplement(num int) int {
	c := 1

	for c < num {
		c <<= 1
		c += 1
	}

	return c ^ num
}
