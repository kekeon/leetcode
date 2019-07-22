package bit_manipulation

func toHex(num int) string {

	if num >= 0 {
		return hex(num)
	}else {
		sum := -4294967296 - num
		return hex(-sum)
	}


}

func hex (num int) string{
	m := map[int]string {
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	var str string
	for num >= 0 {
		if num < 16 {
			str = m[num] + str
			num = 0
			return str
		}else {
			remainder := num % 16
			str = m[remainder] + str
			num = (num - remainder ) / 16
		}
	}

	return str
}