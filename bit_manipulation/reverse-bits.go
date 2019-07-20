package bit_manipulation

func reverseBits(num uint32) uint32 {

	var ans uint32
	for i := 0; i < 32; {

		ans <<= 1
		ans |= num&1
		num >>= 1
		i++
	}

	return ans
}
