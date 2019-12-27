package main

import "fmt"

/**
12. 整数转罗马数字
1994 -> "MCMXCIV"
 */
func intToRoman(num int) string {
	strArr := [][]string{
		{"","I","II","III","IV","V","VI","VII","VIII","IX"},
		{"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"},
		{"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"},
		{"","M","MM","MMM"},
	}

	str := ""
	str = strArr[0][num % 10] + str
	str = strArr[1][num / 10 % 10] + str
	str = strArr[2][num / 100 % 10] + str
	str = strArr[3][num / 1000 % 10] + str
	return str


}

func main() {
	fmt.Println(intToRoman(1994))
}
