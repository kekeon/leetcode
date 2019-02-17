package main

import (
	"fmt"
	"strings"
)

func detectCapitalUse(word string) bool {

	if strings.ToUpper(word) == word {
		return true
	}

	last := strings.ToLower(word[1:])

	if last == word[1:] {
		return true
	}


	return false
}

func main() {
	v := detectCapitalUse("USA")
	fmt.Println(v)
}