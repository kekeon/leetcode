package main

import "fmt"

func toLowerCase(str string) string {

	m := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
		"D": "D",
		"E": "e",
		"F": "f",
		"G": "g",
		"H": "h",
		"I": "i",
		"J": "j",
		"K": "k",
		"L": "l",
		"M": "m",
		"N": "n",
		"O": "o",
		"P": "p",
		"Q": "q",
		"R": "r",
		"S": "s",
		"T": "t",
		"U": "u",
		"V": "v",
		"W": "w",
		"X": "x",
		"Y": "y",
		"Z": "z",
	}

	s := ""
	for i:=0;i<len(str);i++ {

		if m[string(str[i])] != "" {
			s+=m[string(str[i])]
		} else {
		s+=string(str[i])
		}
	}

	return s

}

func main() {
	v := toLowerCase("HelLo")
	fmt.Println(v)
}
