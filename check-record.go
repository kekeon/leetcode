package main

import "fmt"

func checkRecord(s string) bool {

	aCount := 0
	lCount := 0
	for _,v := range s {

		if string(v) == "A" {
			aCount++
		}

		if aCount > 1 {
			return false
		}

		if string(v) == "L" {
			lCount ++
		}else if lCount <= 2{
			lCount = 0
		}
	}

	if aCount <=1 && lCount <=2 {
		return true
	}

	return false
}

func main() {

	v := checkRecord("LLLALL")
	fmt.Println(v)
}
