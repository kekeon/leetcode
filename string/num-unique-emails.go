package main

import (
	"strings"
	"fmt"
)

func numUniqueEmails(emails []string) int {

	num := 0

	if len(emails) == 0 {
		return num
	}

	emailMap := map[string]int{}

	for _,v := range emails {

		end := strings.Index(v,"@")

		frontPrefix := v[:end]
		endPreFix := v[end:]

		addIndex := strings.Index(frontPrefix, "+")

		if addIndex != -1 {
			frontPrefix = frontPrefix[:addIndex]
		}

		frontPrefix = strings.Replace(frontPrefix, ".", "",-1)

		key := frontPrefix + endPreFix

		if emailMap[key] == 0 {
			emailMap[key] = emailMap[key] +1
			num ++
		}
	}

	return num
}

func main() {
	params := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	v := numUniqueEmails(params)
	fmt.Println(v)
}
