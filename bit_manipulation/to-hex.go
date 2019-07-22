package bit_manipulation

import "fmt"

func toHex(num int) string {

	if num < 0 {
		return "ffffffff"
	}

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
		fmt.Println(num)
		if num < 16 {
			str += m[num]
			num = 0
			return str
		}else {

		}
	}

	return str
}
