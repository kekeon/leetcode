package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {

	list := []string{}

	for _,v :=  range strings.Split(s," ") {
		if v != "" {
			list = append(list,v)
		}
	}

	return len(list)
}

func main() {
	v := countSegments("                ")
	fmt.Println(v)
}