package myMath

import (
	"fmt"
	"testing"
)

func TestRecentCounter(t *testing.T) {
	v := mySqrt(36)
	fmt.Println(v)
}

func TestTitleToNumber(t *testing.T) {
	v := titleToNumber2("ZY")
	fmt.Println(v)
}


func TestTrailingZeroes(t *testing.T) {
	v := trailingZeroes(30)
	fmt.Println(v)
}
