package main

import (
	"strings"
	"fmt"
)

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	ss := strings.Replace(s, " ", ",", -1)

	m := []int{}
	for i, v := range ss {
		if string(v) == "," {
			m = append(m,i)
		}
	}
	fmt.Println(m)

	last := ""
	abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(m) == len(s) {
		return 0
	}else if len(m) ==1 {

		start := m[0] +1

		if start == len(s) {
			last = s[0:start]
		}else {
			last = s[start:]
		}

	}else if len(m)==0 && len(s) >0 {
		last = s
	}else {
		for j := len(m) - 1; j>0; j-- {
			start := m[j-1] + 1
			if m[j-1] + 1 != m[j] {
				end := m[j] + 1
				last = s[start:end]
			}

			fmt.Println(last)

		}
	}

	n := strings.Index(abc, string(last[0]))

	if n > 0 {
		return len(last)
	} else {
		return 0
	}

}

func main() {
	s := lengthOfLastWord("A H211 Word")
	fmt.Println("------")
	fmt.Println(s)
}
