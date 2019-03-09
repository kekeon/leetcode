package main

import "fmt"

func buddyStrings(A string, B string) bool {

	if len(A) != len(B) {
		return false
	}

	if len(B) == 2 && B[0] != B[1] &&A[0] == B[0] && A[1] == B[1] {
		return false
	}


	m := map[string]bool{}

	for _, k := range A {
		m[string(k)] = true
	}

	for _, k := range B {
		if !m[string(k)] {
			return false
		}else {
			m[string(k) + "-"] = false
		}
	}


	return true
}

func main() {
	v := buddyStrings("aa", "aa")
	fmt.Println(v)
}
