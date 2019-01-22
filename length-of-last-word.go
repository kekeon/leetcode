package main

import (
	"strings"
	"fmt"
)

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	ss := strings.Replace(s, " ",",",-1)


	fmt.Println(ss)
	arr := strings.Split(ss, "")
	fmt.Println("============")
	fmt.Println(arr)

	filterArr := []string{}
	if len(arr) > 0 {
		for i:=0; i< len(arr); i++ {
			if string(arr[i]) != "," {
				filterArr = append(filterArr, arr[i])
			}
		}
	}

	fmt.Println(filterArr)

	if len(filterArr) == 0 {
		return 0
	}else {
		arr = strings.Split(s, " ")
	}


	abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	last := arr[len(arr) - 1]

	if last == " " {
		return 0
	}

	n := strings.Index(abc, string(last[0]))

	if n > 0 {
		return len(last)
	}else {
		return 0
	}

}

func main() {
	s := lengthOfLastWord("a ")
	fmt.Println("------")
	fmt.Println(s)
}
